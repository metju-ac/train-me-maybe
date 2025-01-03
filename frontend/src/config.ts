const baseUrl = import.meta.env.VITE_API_BASE_URL;
if (!baseUrl || baseUrl === "") {
  throw new Error("VITE_API_BASE_URL is not set");
}

const useMocks = import.meta.env.VITE_USE_MOCKS === "true";

const config = {
  baseUrl,
  useMocks,
};

export default config;
