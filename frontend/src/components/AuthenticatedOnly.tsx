import useIsLoggedIn from "@utils/useIsLoggedIn";
import React from "react";
import { useTranslation } from "react-i18next";

export default function AuthenticatedOnly(props: { element: React.ReactNode }) {
  const isLoggedIn = useIsLoggedIn();
  const { t } = useTranslation("default");

  if (!isLoggedIn) {
    return (
      <>
        <h1>{t("You are not logged in")}</h1>
        <p>{t("To use the app you need to register.")}</p>
      </>
    );
  }

  return props.element;
}
