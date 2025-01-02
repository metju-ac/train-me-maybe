import config from "@/config";
import { getAuthToken } from "@utils/authToken";
import axios from "axios";

const client = axios.create({
  baseURL: config.baseUrl,
  timeout: 2000,
});

client.interceptors.request.use((config) => {
  const token = getAuthToken();
  if (token) {
    config.headers["Authorization"] = `Bearer ${token}`;
  }
  return config;
});

export default client;
