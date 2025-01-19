let baseUrl = import.meta.env.VITE_API_BASE_URL;
if (!baseUrl || baseUrl === "") {
  baseUrl = window.location.origin + "/api";
}

const useMocks = import.meta.env.VITE_USE_MOCKS === "true";

const supportEmail = import.meta.env.VITE_SUPPORT_EMAIL;
if (!supportEmail || supportEmail === "") {
  throw new Error("VITE_SUPPORT_EMAIL is not set");
}

const config = {
  baseUrl,
  useMocks,
  supportEmail,
};

export default config;
