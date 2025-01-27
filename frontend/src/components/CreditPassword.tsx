import Visibility from "@mui/icons-material/Visibility";
import VisibilityOff from "@mui/icons-material/VisibilityOff";
import {
  IconButton,
  InputAdornment,
  TextField,
  TextFieldProps,
  Tooltip,
} from "@mui/material";
import { useState } from "react";
import { useTranslation } from "react-i18next";

export default function CreditPassword({
  creditPassword,
  setCreditPassword,
  ...props
}: {
  creditPassword: string;
  setCreditPassword: (creditPassword: string) => void;
} & TextFieldProps) {
  const [showPassword, setShowPassword] = useState(false);

  const handleClickShowPassword = () => {
    setShowPassword(!showPassword);
  };

  const handleMouseDownPassword = (
    event: React.MouseEvent<HTMLButtonElement>
  ) => {
    event.preventDefault();
  };
  const { t } = useTranslation("default");

  return (
    <TextField
      label={t("Credit Password")}
      type={showPassword ? "text" : "password"}
      value={creditPassword}
      onChange={(e) => {
        setCreditPassword(e.target.value);
      }}
      fullWidth
      slotProps={{
        input: {
          endAdornment: (
            <InputAdornment position="end">
              <Tooltip title={t("Toggle password visibility")}>
                <IconButton
                  aria-label={t("Toggle password visibility")}
                  onClick={handleClickShowPassword}
                  onMouseDown={handleMouseDownPassword}
                  edge="end"
                >
                  {showPassword ? <VisibilityOff /> : <Visibility />}
                </IconButton>
              </Tooltip>
            </InputAdornment>
          ),
        },
      }}
      {...props}
    />
  );
}
