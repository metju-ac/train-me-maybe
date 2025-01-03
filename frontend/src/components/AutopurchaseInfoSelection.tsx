import { User } from "@models/User";
import {
  Box,
  Button,
  Checkbox,
  CircularProgress,
  FormControl,
  FormControlLabel,
  TextField,
  Tooltip,
  Typography,
} from "@mui/material";
import useUpdateUserDetails from "@utils/useUpdateUserDetails";
import useUserDetails from "@utils/useUserDetails";
import React, { useContext, useState } from "react";
import CreditPassword from "./CreditPassword";
import { ToastBarContext } from "./ToastBarProvider";

interface AutopurchaseInfoSelectionProps {
  handleSubmit: (ev: React.FormEvent) => void;
}

export default function AutopurchaseInfoSelection(
  props: AutopurchaseInfoSelectionProps
) {
  const { data, isLoading, isError } = useUserDetails();

  if (isLoading) {
    return <CircularProgress />;
  }

  if (isError) {
    return <div>Error while loading data</div>;
  }

  return <AutopurchaseInfoSelectionForm {...props} user={data!} />;
}

function AutopurchaseInfoSelectionForm({
  user,
  ...props
}: AutopurchaseInfoSelectionProps & { user: User }) {
  const [creditUser, setCreditUser] = useState(user.creditUser || "");
  const [creditPassword, setCreditPassword] = useState(
    user.creditPassword || ""
  );
  const [cutOffTime, setCutOffTime] = useState(user.cutOffTime || "");
  const [minimalCredit, setMinimalCredit] = useState(
    `${user.minimalCredit}` || ""
  );

  const [saveInfoToAccount, setSaveInfoToAccount] = useState(false);

  const updateUserDetails = useUpdateUserDetails();
  const { setToast } = useContext(ToastBarContext);

  const handleSubmit = (e: React.FormEvent) => {
    if (saveInfoToAccount) {
      return updateUserDetails.mutate(
        {
          creditUser,
          creditPassword,
          cutOffTime,
          minimalCredit: Number(minimalCredit),
        },
        {
          onSuccess: () => props.handleSubmit(e),
          onError: (err) => {
            setToast(
              "Failed to save details to account: " + err.message,
              "error"
            );
          },
        }
      );
    }

    props.handleSubmit(e);
  };

  const isValid = creditUser !== "" && creditPassword !== "";

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
      <Box>
        <Typography variant="h5">Review purchase options</Typography>
      </Box>
      <TextField
        label="Credit Number"
        required
        error={creditUser === ""}
        value={creditUser}
        onChange={(e) => setCreditUser(e.target.value)}
        fullWidth
        margin="normal"
      />

      <CreditPassword
        creditPassword={creditPassword}
        setCreditPassword={setCreditPassword}
        error={creditPassword === ""}
        required
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

      <FormControl fullWidth>
        <FormControlLabel
          control={
            <Checkbox
              checked={saveInfoToAccount}
              onChange={(e) => setSaveInfoToAccount(e.target.checked)}
              inputProps={{ "aria-label": "controlled" }}
            />
          }
          label="Save info to account?"
        />
      </FormControl>

      <Tooltip title={isValid ? "" : "First fill out all necessary info"}>
        <span>
          <Button
            type="submit"
            variant="contained"
            color="primary"
            disabled={!isValid}
          >
            Start watching route
          </Button>
        </span>
      </Tooltip>
    </Box>
  );
}
