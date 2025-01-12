import { useTranslation } from "react-i18next";

export default function NotFound() {
  const { t } = useTranslation("default");

  return (
    <>
      <h1>{t("404 Not Found")}</h1>
      <p>{t("The page you want to see is not there 😢.")}</p>
    </>
  );
}
