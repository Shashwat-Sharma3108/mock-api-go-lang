import { handleRemoveData } from '@/apis/uploadconfig';
import React, { useState } from 'react';
import { BASE_URL } from '@/constants/apiEndpoints';

const ApiPage = ({ apis, setApis }) => {
  const [expandedApiId, setExpandedApiId] = useState(null);

  const getMethodColor = (method) => {
    switch (method) {
      case 'GET':
        return 'bg-green-200 text-green-800';
      case 'POST':
        return 'bg-blue-500 text-white';
      case 'PUT':
      case 'PATCH':
        return 'bg-blue-200 text-blue-800';
      case 'DELETE':
        return 'bg-red-200 text-red-800';
      default:
        return 'bg-gray-200 text-gray-800';
    }
  };

  const toggleExpand = (id) => {
    setExpandedApiId(expandedApiId === id ? null : id);
  };

  const generateCurl = (api) => {
    const method = api.method;
    const url = `${BASE_URL}${api.url}`;  // Ensure full URL
    const headers = {
      'Content-Type': 'application/json',
    };

    let curlCommand = `curl -X ${method} '${url}'`;
    
    // Add headers
    for (let header in headers) {
      curlCommand += ` -H '${header}: ${headers[header]}'`;
    }

    // Add request body for POST, PUT, PATCH
    if ((method === 'POST' || method === 'PUT' || method === 'PATCH') && api.request) {
      curlCommand += ` -d '${JSON.stringify(api.request)}'`;  // Fix JSON.stringify issue
    }

    return curlCommand;
};


  const deleteEndPoint = async (id) => {
    try {
      const response = await handleRemoveData(id);
      const updatedApis = apis.filter(a => a?.id !== id);
      setApis(updatedApis)
    } catch (error) {
      console.error({ error })
    }
  };

  const copyToClipboard = (curlCommand) => {
    navigator.clipboard.writeText(curlCommand).then(() => {
      alert('cURL command copied to clipboard!');
    });
  };

  return (
    <section className="w-2/3 p-6 bg-white overflow-y-auto" style={{ maxHeight: '100vh', position: 'relative' }}>
      <h1 className="text-3xl font-bold text-gray-900 mb-4">API Documentation</h1>
      <div className="space-y-6">
        {apis?.length > 0 ? (
          apis?.map((api) => (
            <div key={api.id} className="p-6 border border-gray-300 rounded-lg shadow-lg">
              <h2 className="text-2xl font-semibold text-gray-800 mb-2">
                <button
                  className="flex items-center justify-between w-full text-left"
                  onClick={() => toggleExpand(api.id)}
                >
                  API {api.id}: {api.url}
                  <span className="text-sm text-gray-600">
                    {expandedApiId === api.id ? 'Hide' : 'Show'} Details
                  </span>
                </button>
              </h2>

              <div className={`text-lg font-medium p-2 rounded-lg mb-4 ${getMethodColor(api.method)}`}>
                <strong>Request Type:</strong> {api.method}
              </div>

              {expandedApiId === api.id && (
                <div className="bg-gray-100 p-4 rounded-lg">
                  <h3 className="text-xl font-medium text-gray-800 mb-2">Request</h3>
                  <pre className="bg-gray-800 text-white p-4 rounded-lg mb-4">{JSON.stringify(api?.request || {}, null, 2)}</pre>

                  <h3 className="text-xl font-medium text-gray-800 mb-2">Response</h3>
                  <pre className="bg-gray-800 text-white p-4 rounded-lg mb-4">{JSON.stringify(api?.response, null, 2)}</pre>

                  {api?.headers && (
                    <>
                      <h3 className="text-xl font-medium text-gray-800 mb-2">Headers</h3>
                      <pre className="bg-gray-800 text-white p-4 rounded-lg mb-4">{JSON.stringify(api?.headers, null, 2)}</pre>
                    </>
                  )}

                  <button
                    onClick={() => copyToClipboard(generateCurl(api))}
                    className="mt-4 bg-gray-800 text-white p-2 rounded-lg w-full"
                  >
                    Copy cURL Command
                  </button>
                  <button
                    onClick={() => deleteEndPoint(api?.id)}
                    className="mt-4 bg-red-600 text-white p-2 rounded-lg w-full"
                  >
                    Delete API
                  </button>
                </div>
              )}
            </div>
          ))
        ) : (
          <p className="text-gray-600">No APIs available. Please upload a valid YAML file to see APIs.</p>
        )}
      </div>
    </section>
  );
};

export default ApiPage;
