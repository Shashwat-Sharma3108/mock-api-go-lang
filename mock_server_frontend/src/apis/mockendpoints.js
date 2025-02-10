import { API_END_POINT } from '@/constants/apiEndpoints'
import apiClient, { handleApiError } from './index'

export const handleUploadData = async (postData) => {
    try {
        const res = await apiClient.post(API_END_POINT.UPLOAD_CONFIG, postData)
        return res;
    } catch (error) {
        handleApiError(error);
        throw error;
    }
}

