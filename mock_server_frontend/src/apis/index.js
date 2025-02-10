import axios from 'axios';
import { BASE_URL } from '@/constants/apiEndpoints';

const apiClient = axios.create({
    baseURL: BASE_URL,
    headers: { 
        'Content-Type': 'application/json',
      },
      withCredentials: false,  // If your server supports credentials (cookies, sessions, etc.)
    
});


export const handleApiError = (error) => {
    console.error('API Error:', error.response ? error.response.data : error.message);
    alert('An error occurred while processing the request. Please try again.');
};


export default apiClient;