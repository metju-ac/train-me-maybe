import axios from "axios";
import config from "@/config";

const client = axios.create({
  baseURL: config.baseUrl,
  timeout: 2000,
});

export default client;
