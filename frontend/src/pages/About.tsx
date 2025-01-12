import { useTranslation } from "react-i18next";

export default function About() {
  const { t } = useTranslation("default");

  return (
    <>
      <h1>{t("About")}</h1>
      <p>{t("Train me maybe about detailed description")}</p>
      <img
        src="/regiojet-large.jpg"
        alt="RegioJet"
        style={{
          maxWidth: "100%",
        }}
      />
    </>
  );
}
