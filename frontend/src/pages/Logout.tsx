import { removeAuthToken } from "@utils/authToken";
import { useEffect } from "react";
import { useTranslation } from "react-i18next";

export default function Logout() {
  useEffect(() => {
    removeAuthToken();
  }, []);

  const { t } = useTranslation("default");

  return (
    <>
      <h1>{t("Logout")}</h1>
      <p>{t("You have been logged out.")}</p>
    </>
  );
}
