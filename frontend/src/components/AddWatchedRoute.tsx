import AddCircleIcon from "@mui/icons-material/AddCircle";
import { IconButton, Stack, Tooltip } from "@mui/material";
import { useTranslation } from "react-i18next";
import { useNavigate } from "react-router";

export default function AddWatchedRoute() {
  const navigate = useNavigate();
  const { t } = useTranslation("default");

  return (
    <Stack direction="row" justifyContent="flex-end">
      <Tooltip arrow title={t("Add a new watched route")}>
        <IconButton onClick={() => void navigate("/routes/new")}>
          <AddCircleIcon color="primary" sx={{ fontSize: "3rem" }} />
        </IconButton>
      </Tooltip>
    </Stack>
  );
}
