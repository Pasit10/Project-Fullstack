interface ImportMetaEnv {
  readonly API_KEY: string;
  readonly AUTH_DOMAIN: string;
  readonly PROJECT_ID: string;
  readonly STORAGE_BUCKET: string;
  readonly MESSAGINGSENDER_ID: string;
  readonly APP_ID: string;
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}
