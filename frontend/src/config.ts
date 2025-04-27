let baseUrl = import.meta.env.VITE_API_BASE_URL;
if (!baseUrl || baseUrl === "") {
  baseUrl = window.location.origin + "/api";
}

const useMocks = import.meta.env.VITE_USE_MOCKS === "true";

const supportEmail = import.meta.env.VITE_SUPPORT_EMAIL;
if (!supportEmail || supportEmail === "") {
  throw new Error("VITE_SUPPORT_EMAIL is not set");
}

let urlPrefix = import.meta.env.VITE_URL_PREFIX;
if (typeof urlPrefix !== "string") {
  urlPrefix = "";
}
if (urlPrefix.startsWith("/")) {
  urlPrefix = urlPrefix.slice(1);
}
if (urlPrefix.endsWith("/")) {
  urlPrefix = urlPrefix.slice(0, -1);
}

const config = {
  baseUrl,
  useMocks,
  supportEmail,
  urlPrefix,
};

export default config;
