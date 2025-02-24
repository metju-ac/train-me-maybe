import config from "@/config";
import { getAuthToken, removeAuthToken } from "@utils/authToken";
import axios from "axios";

const client = axios.create({
  baseURL: config.baseUrl,
  timeout: 2000,
});

client.interceptors.request.use((config) => {
  const token = getAuthToken();
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

client.interceptors.response.use(
    response => response,
    error => {
        if (error.response && error.response.status === 401) {
            removeAuthToken();
            if (typeof window !== "undefined") {
                window.location.href = "/login";
            }
        }
        return Promise.reject(error);
    }
);

export default client;
