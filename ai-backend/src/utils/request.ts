import axios from 'axios';
import {BASE_URL} from "../apis/url";

export const request = ({url, method, params}: { url: string; method: string; params?: any }) => {
    return axios({
        headers: {},
        baseURL: BASE_URL,
        url: url,
        method: method,
        data: method !== 'GET' ? params : undefined,
        params: method === 'GET' ? params : undefined,
    }).then(response => {
        return response.data;
    }).catch(error => {
        console.error('There was an error with the request:', error);
    });
};

