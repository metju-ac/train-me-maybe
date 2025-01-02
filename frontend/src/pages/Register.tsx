import { ToastBarContext } from "@components/ToastBarProvider";
import { Box, Button, TextField, Typography } from "@mui/material";
import userService from "@services/userService";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import React, { useContext, useState } from "react";
import { useNavigate } from "react-router";

export default function Register() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const queryClient = useQueryClient();

  // Mutations
  const mutation = useMutation({
    mutationKey: [userService.registerUser.key],
    mutationFn: userService.registerUser.fn,
    onSuccess: () => {
      // Invalidate and refetch
      queryClient.invalidateQueries({
        queryKey: [userService.getUserDetails.key],
      });
    },
  });

  const { setToast } = useContext(ToastBarContext);
  let navigate = useNavigate();

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();
    // Handle form submission logic here
    console.log("Email:", email);
    console.log("Password:", password);

    mutation.mutate(
      { email, password },
      {
        onSuccess: (data) => {
          console.log("User registered successfully", data);
          setToast("Successfully registered user", "success");
          navigate("/");
        },
        onError: (error) => {
          console.error("Error registering user", error);
          setToast("Error registering user: " + error.message, "error");
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
        Register
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
        Register
      </Button>
    </Box>
  );
}
