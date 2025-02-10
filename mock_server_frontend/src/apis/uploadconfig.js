import { API_END_POINT } from '@/constants/apiEndpoints'
import apiClient, { handleApiError } from './index'

export const getUploadedEndpoint = async () => {
    try {
        const res = await apiClient.get(API_END_POINT.GET_END_POINTS)
        return res;
    } catch (error) {
        handleApiError(error);
        throw error;
    }
}

export const handleRemoveData = async (id) => {
    try {
        const res = await apiClient.delete(`${API_END_POINT.DELETE_ENDPOINTS}/${id}`)
        return res;
    } catch (error) {
        handleApiError(error);
        throw error;
    }
}