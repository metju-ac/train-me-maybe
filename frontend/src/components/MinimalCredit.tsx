import { Help } from "@mui/icons-material";
import {
  InputAdornment,
  TextField,
  TextFieldProps,
  Tooltip,
} from "@mui/material";

export default function MinimalCredit({
  minimalCredit,
  setMinimalCredit,
  ...props
}: {
  minimalCredit: string;
  setMinimalCredit: (minimalCredit: string) => void;
} & TextFieldProps) {
  return (
    <TextField
      label="Minimal Credit"
      type="number"
      value={minimalCredit}
      onChange={(e) => setMinimalCredit(e.target.value)}
      fullWidth
      slotProps={{
        input: {
          endAdornment: (
            <InputAdornment position="end">
              <Tooltip
                arrow
                title="The minimal credit you want to have on your account. If your credit falls below this value, you will be notified by email."
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
