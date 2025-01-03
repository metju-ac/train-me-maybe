import { Help } from "@mui/icons-material";
import {
  InputAdornment,
  TextField,
  TextFieldProps,
  Tooltip,
} from "@mui/material";

export default function CutOffTime({
  cutOffTime,
  setCutOffTime,
  ...props
}: {
  cutOffTime: string;
  setCutOffTime: (cutOffTime: string) => void;
} & TextFieldProps) {
  return (
    <TextField
      label="Cut Off Time"
      type="number"
      value={cutOffTime}
      onChange={(e) => setCutOffTime(e.target.value)}
      fullWidth
      slotProps={{
        input: {
          endAdornment: (
            <InputAdornment position="end">
              <Tooltip title="Time in minutes before the train departure when the last autopurchase can be made. E.g. 120 means that watched routes will be purchased no less than 2 hours before the train departure.">
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