import { Help } from "@mui/icons-material";
import {
  InputAdornment,
  TextField,
  TextFieldProps,
  Tooltip,
} from "@mui/material";
import { useTranslation } from "react-i18next";

export default function CreditNumber({
  creditUser,
  setCreditUser,
  ...props
}: {
  creditUser: string;
  setCreditUser: (creditUser: string) => void;
} & TextFieldProps) {
  const { t } = useTranslation("default");

  return (
    <TextField
      label={t("Credit Number")}
      value={creditUser}
      onChange={(e) => {
        setCreditUser(e.target.value);
      }}
      fullWidth
      slotProps={{
        input: {
          endAdornment: (
            <InputAdornment position="end">
              <Tooltip
                arrow
                title={t(
                  "The Regiojet number of your credit ticket (e.g. 123456789). A.k.a. 'číslo kreditové jizdenky'"
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
