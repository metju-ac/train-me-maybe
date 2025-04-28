import config from "@/config";
import { Link } from "@mui/material";
import { useTranslation } from "react-i18next";

export default function About() {
  const { t } = useTranslation("default");

  return (
    <>
      <h1>{t("About")}</h1>
      <p>{t("Train me maybe about detailed description paragraph 1")}</p>
      <p>
        {t("Train me maybe about detailed description paragraph 2")}
        <Link href={`mailto:${config.supportEmail}`}>
          {config.supportEmail}
        </Link>
        .
      </p>
      <p>
        {t("Train me maybe about detailed description paragraph 3")}
        <Link
          href="https://github.com/metju-ac/train-me-maybe/"
          target="_blank"
        >
          github.com/metju-ac/train-me-maybe
        </Link>
      </p>
      <img
        src={`/${config.urlPrefix}/img/regiojet-large.jpg`}
        alt="RegioJet"
        title={t("Our app finds free seats with flying colors!")}
        style={{
          maxWidth: "100%",
        }}
      />
    </>
  );
}
