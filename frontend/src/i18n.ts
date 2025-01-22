import i18n from "i18next";
import LanguageDetector from "i18next-browser-languagedetector";
import { initReactI18next } from "react-i18next";
import config from "./config";
import locales from "./locales";

// just an alias
export const resources = locales;

export const availableLanguages = Object.keys(
  locales
) as (keyof typeof locales)[];
export type AvailableLanguage = (typeof availableLanguages)[number];

export const languageToFlagUrl: Record<AvailableLanguage, string> = {
  cs: "https://flagcdn.com/w20/cz.png",
  en: "https://flagcdn.com/w20/us.png",
  sk: "https://flagcdn.com/w20/sk.png",
};

export const initI18N = () => {
  return i18n
    .use(LanguageDetector)
    .use(initReactI18next)
    .init({
      debug: config.useMocks,
      supportedLngs: availableLanguages,
      resources,
      fallbackLng: "en",

      interpolation: {
        escapeValue: false,
      },
    });
};
