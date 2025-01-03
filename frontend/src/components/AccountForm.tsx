import { User } from "@models/User";
import { Box, Button, TextField, Typography } from "@mui/material";
import useUpdateUserDetails from "@utils/useUpdateUserDetails";
import React, { useState } from "react";
import CreditPassword from "./CreditPassword";

export default function AccountForm({ user }: { user: User }) {
  const [creditUser, setcreditUser] = useState(user.creditUser || "");
  const [creditPassword, setCreditPassword] = useState(
    user.creditPassword || ""
  );
  const [cutOffTime, setCutOffTime] = useState(user.cutOffTime || "");
  const [minimalCredit, setMinimalCredit] = useState(
    `${user.minimalCredit}` || ""
  );

  const mutation = useUpdateUserDetails();

  const handleSubmit = (event: React.FormEvent) => {
    event.preventDefault();
    mutation.mutate({
      creditUser,
      creditPassword,
      cutOffTime,
      minimalCredit: Number(minimalCredit),
    });
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
      <CreditPassword
        creditPassword={creditPassword}
        setCreditPassword={setCreditPassword}
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
