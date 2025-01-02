import { User } from "@models/User";
import Visibility from "@mui/icons-material/Visibility";
import VisibilityOff from "@mui/icons-material/VisibilityOff";
import {
  Box,
  Button,
  IconButton,
  InputAdornment,
  TextField,
  Typography,
} from "@mui/material";
import userService from "@services/userService";
import { useMutation } from "@tanstack/react-query";
import React, { useContext, useState } from "react";
import { ToastBarContext } from "./ToastBarProvider";

export default function AccountForm({ user }: { user: User }) {
  const [creditUser, setcreditUser] = useState(user.creditUser || "");
  const [creditPassword, setCreditPassword] = useState(
    user.creditPassword || ""
  );
  const [cutOffTime, setCutOffTime] = useState(user.cutOffTime || "");
  const [minimalCredit, setMinimalCredit] = useState(
    `${user.minimalCredit}` || ""
  );

  const { setToast } = useContext(ToastBarContext);

  const mutation = useMutation({
    mutationKey: [userService.updateUserDetails.key],
    mutationFn: userService.updateUserDetails.fn,
    onSuccess: () => {
      setToast("Details updated successfully!", "success");
    },
    onError: () => {
      setToast("Failed to update details!", "error");
    },
  });

  const handleSubmit = (event: React.FormEvent) => {
    event.preventDefault();
    mutation.mutate({
      creditUser,
      creditPassword,
      cutOffTime,
      minimalCredit: Number(minimalCredit),
    });
  };

  const [showPassword, setShowPassword] = useState(false);

  const handleClickShowPassword = () => {
    setShowPassword(!showPassword);
  };

  const handleMouseDownPassword = (
    event: React.MouseEvent<HTMLButtonElement>
  ) => {
    event.preventDefault();
  };

  return (
    <Box
      component="form"
      onSubmit={handleSubmit}
      sx={{
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
        justifyContent: "center",
        padding: 2,
      }}
    >
      <Typography variant="h4" component="h1" gutterBottom>
        Account Details
      </Typography>
      <Typography variant="h6" component="h2" gutterBottom>
        Email: {user.email}
      </Typography>
      <TextField
        label="Credit Number"
        value={creditUser}
        onChange={(e) => setcreditUser(e.target.value)}
        fullWidth
        margin="normal"
      />
      <TextField
        label="Credit Password"
        type={showPassword ? "text" : "password"}
        value={creditPassword}
        onChange={(e) => setCreditPassword(e.target.value)}
        fullWidth
        margin="normal"
        InputProps={{
          endAdornment: (
            <InputAdornment position="end">
              <IconButton
                aria-label="toggle password visibility"
                onClick={handleClickShowPassword}
                onMouseDown={handleMouseDownPassword}
                edge="end"
              >
                {showPassword ? <VisibilityOff /> : <Visibility />}
              </IconButton>
            </InputAdornment>
          ),
        }}
      />
      <TextField
        label="Cut Off Time"
        type="number"
        value={cutOffTime}
        onChange={(e) => setCutOffTime(e.target.value)}
        fullWidth
        margin="normal"
      />
      <TextField
        label="Minimal Credit"
        type="number"
        value={minimalCredit}
        onChange={(e) => setMinimalCredit(e.target.value)}
        fullWidth
        margin="normal"
      />
      <Button type="submit" variant="contained" color="primary" sx={{ mt: 2 }}>
        Update Details
      </Button>
    </Box>
  );
}
