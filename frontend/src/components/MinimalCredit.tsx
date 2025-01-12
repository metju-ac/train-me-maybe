import { Help } from "@mui/icons-material";
import {
  InputAdornment,
  TextField,
  TextFieldProps,
  Tooltip,
} from "@mui/material";
import { useTranslation } from "react-i18next";

export default function MinimalCredit({
  minimalCredit,
  setMinimalCredit,
  ...props
}: {
  minimalCredit: string;
  setMinimalCredit: (minimalCredit: string) => void;
} & TextFieldProps) {
  const { t } = useTranslation("default");

  return (
    <TextField
      label={t("Minimal Credit")}
      type="number"
      value={minimalCredit}
      onChange={(e) => {
        setMinimalCredit(e.target.value);
      }}
      fullWidth
      slotProps={{
        input: {
          endAdornment: (
            <InputAdornment position="end">
              <Tooltip
                arrow
                title={t(
                  "The minimal credit you want to have on your account. If your credit falls below this value, you will be notified by email."
                )}
              >
                <Help />
              </Tooltip>
            </InputAdornment>
          ),
        },
      }}
      {...props}
    />
  );
}
