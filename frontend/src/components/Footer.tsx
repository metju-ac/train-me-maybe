import config from "@/config";
import { Box, Link } from "@mui/material";
import { useTranslation } from "react-i18next";

export default function Footer() {
  const { t } = useTranslation("default");

  return (
    <Box marginY={2}>
      {t("Robert Gemrot & Matěj Klíma, 2025. Contact us at ")}
      <Link href={`mailto:${config.supportEmail}`}>{config.supportEmail}</Link>.
    </Box>
  );
}
