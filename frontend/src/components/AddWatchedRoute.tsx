import AddCircleIcon from "@mui/icons-material/AddCircle";
import { IconButton, Stack, Tooltip } from "@mui/material";
import { useNavigate } from "react-router";

export default function AddWatchedRoute() {
  const navigate = useNavigate();

  return (
    <Stack direction="row" justifyContent="flex-end">
      <Tooltip arrow title="Add a new watched route">
        <IconButton onClick={() => navigate("/routes/new")}>
          <AddCircleIcon color="primary" sx={{ fontSize: "3rem" }} />
        </IconButton>
      </Tooltip>
    </Stack>
  );
}
