import config from "@/config";
import { Box, Link } from "@mui/material";

export default function Footer() {
  return (
    <Box marginY={2}>
      Robert Gemrot & Matěj Klíma, 2025. Contact us at{" "}
      <Link href={`mailto:${config.supportEmail}`}>{config.supportEmail}</Link>.
    </Box>
  );
}
