import SubmitButton from "@components/SubmitButton";
import { Box, TextField, Typography } from "@mui/material";
import useLoginUser from "@utils/useLoginUser";
import React, { useState } from "react";

export default function Login() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  // Mutations
  const mutation = useLoginUser();

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();

    mutation.mutate({ email, password });
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
        gap: 2,
      }}
    >
      <Typography variant="h4" component="h1" gutterBottom>
        Login
      </Typography>
      <TextField
        label="Email"
        type="email"
        value={email}
        onChange={(e) => { setEmail(e.target.value); }}
        required
        fullWidth
      />
      <TextField
        label="Password"
        type="password"
        value={password}
        onChange={(e) => { setPassword(e.target.value); }}
        required
        fullWidth
      />
      <SubmitButton
        isValid={email !== "" && password !== ""}
        tooltipTitle="First fill out the form"
      >
        Login
      </SubmitButton>
    </Box>
  );
}
