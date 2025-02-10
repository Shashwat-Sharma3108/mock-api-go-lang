
import UploadSection from '@/components/UploadSection';
import ApiPage from '@/components/ApiPage';
import React, { useEffect, useState } from 'react';
import { handleUploadData } from '@/apis/mockendpoints';
import { getUploadedEndpoint } from '@/apis/uploadconfig';

const Home = () => {
  const [apis, setApis] = useState([])



  const getTheListOfEndPoints = async () => {
    try {
      const uploadedEndPointResponse = await getUploadedEndpoint();
      return uploadedEndPointResponse;
    } catch (error) {
      throw error;
    }
  }

  const onDataUpload = async (data) => {
    try {
      const response = await handleUploadData(data);
      if (response) {
        const uploadedEndPointResponse = await getTheListOfEndPoints();
        setApis(uploadedEndPointResponse?.data || [])
      }
    } catch (error) {
      console.error({ error })
    }

  };

  const handleGetAllEndPoints = async () => {
    try {
      const response = await getTheListOfEndPoints();
      setApis(response?.data || []);
    } catch (error) {
      console.error({ error })
    }
  }

  useEffect(() => {
    handleGetAllEndPoints();
  }, [])


  return (
    <>
      <div className="main-container flex">
        <UploadSection onDataUpload={onDataUpload} className="upload-section" />
        <ApiPage apis={apis} setApis={setApis} className="api-section" />
      </div>
    </>
  );
}

export default Home;
