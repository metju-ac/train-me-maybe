import { Help } from "@mui/icons-material";
import {
  InputAdornment,
  TextField,
  TextFieldProps,
  Tooltip,
} from "@mui/material";

export default function CreditNumber({
  creditUser,
  setCreditUser,
  ...props
}: {
  creditUser: string;
  setCreditUser: (creditUser: string) => void;
} & TextFieldProps) {
  return (
    <TextField
      label="Credit Number"
      value={creditUser}
      onChange={(e) => { setCreditUser(e.target.value); }}
      fullWidth
      slotProps={{
        input: {
          endAdornment: (
            <InputAdornment position="end">
              <Tooltip
                arrow
                title="The Regiojet number of your credit ticket (e.g. 123456789). A.k.a. 'číslo kreditové jizdenky'"
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
