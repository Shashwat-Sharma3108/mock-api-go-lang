import React, { useState } from 'react';
import ReactJson from 'react-json-view';

const UploadSection = ({ onDataUpload }) => {
  const [jsonData, setJsonData] = useState(null);

  const handleFileUpload = (event) => {
    const file = event.target.files[0];
    if (!file) return;

    const reader = new FileReader();
    reader.onload = () => {
      try {
        if (file.type === 'application/json' || file.name.endsWith('.json')) {
          const data = JSON.parse(reader.result);
          setJsonData(data);
          onDataUpload(data);
        } else {
          alert('Please upload a valid JSON file');
        }
      } catch (error) {
        console.error('Error parsing file:', error);
      }
    };
    reader.readAsText(file);
  };

  // const handleUpdateJSONData = (data) => {
  //   setJsonData(data);
  //   onDataUpload(data);
  // }

  return (
    <section className="w-1/3 p-6 bg-white border-r border-gray-300">
      <h2 className="text-xl font-semibold text-gray-800 mb-4">Upload JSON File</h2>
      <input
        type="file"
        accept=".json"
        onChange={handleFileUpload}
        className="bg-gray-100 p-2 border border-gray-300 rounded-lg mb-4"
      />
      <p className="text-gray-600">Upload a JSON file to view and edit data</p>

      {/* {jsonData && (
        <div className="mt-4">
          <h3 className="text-lg font-semibold text-gray-800">Edit JSON Data</h3>
          <ReactJson
            src={jsonData}
            onEdit={(edit) => handleUpdateJSONData(edit?.updated_src)}
            onAdd={(add) => handleUpdateJSONData(add?.updated_src)}
            onDelete={(del) => handleUpdateJSONData(del?.updated_src)}
            theme="monokai"
            displayDataTypes={false}
          />
        </div>
      )} */}
    </section>
  );
};

export default UploadSection;