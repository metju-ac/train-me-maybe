import { ToastBarContext } from "@components/ToastBarProvider";
import { Box, Button, TextField, Typography } from "@mui/material";
import userService from "@services/userService";
import { useMutation } from "@tanstack/react-query";
import React, { useContext, useState } from "react";
import { useNavigate } from "react-router";

export default function Login() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  // Mutations
  const mutation = useMutation({
    mutationKey: [userService.loginUser.key],
    mutationFn: userService.loginUser.fn,
  });

  const { setToast } = useContext(ToastBarContext);
  let navigate = useNavigate();

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();

    mutation.mutate(
      { email, password },
      {
        onSuccess: (data) => {
          console.log("User logged in successfully", data);
          setToast("Successfully logged in", "success");
          navigate("/account");
        },
        onError: (error) => {
          console.error("Error logging in", error);
          setToast("Error logging in: " + error.message, "error");
        },
      }
    );
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
        Login
      </Typography>
      <TextField
        label="Email"
        type="email"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
        required
        fullWidth
        margin="normal"
      />
      <TextField
        label="Password"
        type="password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
        required
        fullWidth
        margin="normal"
      />
      <Button type="submit" variant="contained" color="primary" sx={{ mt: 2 }}>
        Login
      </Button>
    </Box>
  );
}
