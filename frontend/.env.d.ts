interface ImportMetaEnv {
  readonly VITE_API_BASE_URL: string;
  readonly VITE_USE_MOCKS: string | undefined;
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}
