interface ImportMetaEnv {
  readonly VITE_API_BASE_URL: string | undefined;
  readonly VITE_USE_MOCKS: string | undefined;
  readonly VITE_SUPPORT_EMAIL: string | undefined;
  readonly VITE_URL_PREFIX: string | undefined;
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}
