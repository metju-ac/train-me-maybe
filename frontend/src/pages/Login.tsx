import SubmitButton from "@components/SubmitButton";
import { Box, TextField, Typography } from "@mui/material";
import useLoginUser from "@utils/useLoginUser";
import React, { useState } from "react";
import { useTranslation } from "react-i18next";

export default function Login() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const { t } = useTranslation("default");

  // Mutations
  const mutation = useLoginUser();

  const handleSubmit = (event: React.FormEvent) => {
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
        {t("Login")}
      </Typography>
      <TextField
        label={t("Email")}
        type="email"
        value={email}
        onChange={(e) => {
          setEmail(e.target.value);
        }}
        required
        fullWidth
      />
      <TextField
        label={t("Password")}
        type="password"
        value={password}
        onChange={(e) => {
          setPassword(e.target.value);
        }}
        required
        fullWidth
      />
      <SubmitButton
        isValid={email !== "" && password !== ""}
        tooltipTitle={t("First fill out the form")}
      >
        {t("Login")}
      </SubmitButton>
    </Box>
  );
}
