import { User } from "@models/User";
import {
  Box,
  Checkbox,
  CircularProgress,
  FormControl,
  FormControlLabel,
  Typography,
} from "@mui/material";
import { Tariff } from "@services/staticDataService";
import printError from "@utils/printError";
import useTariffs from "@utils/useTariffs";
import useUpdateUserDetails from "@utils/useUpdateUserDetails";
import useUserDetails from "@utils/useUserDetails";
import React, { useContext, useState } from "react";
import CreditNumber from "./CreditNumber";
import CreditPassword from "./CreditPassword";
import CutOffTime from "./CutOffTime";
import MinimalCredit from "./MinimalCredit";
import SelectTariff from "./SelectTariff";
import SubmitButton from "./SubmitButton";
import ToastBarContext from "./ToastBarContext";

export interface AutopurchaseInfoSelectionProps {
  handleSubmit: (
    ev: React.FormEvent,
    data: {
      creditUser: string;
      creditPassword: string;
      cutOffTime?: number;
      minimalCredit?: number;
      tariffKey: string;
    }
  ) => void;
}

export default function AutopurchaseInfoSelection(
  props: AutopurchaseInfoSelectionProps
) {
  const { data, isLoading, isError } = useUserDetails();
  const {
    data: tariffs,
    isLoading: isLoadingTariffs,
    isError: isErrorTariffs,
  } = useTariffs();

  if (isLoading || isLoadingTariffs) {
    return <CircularProgress />;
  }

  if (isError || isErrorTariffs) {
    return <div>Error while loading data</div>;
  }

  return (
    <AutopurchaseInfoSelectionForm {...props} user={data!} tariffs={tariffs!} />
  );
}

function AutopurchaseInfoSelectionForm({
  user,
  tariffs,
  ...props
}: AutopurchaseInfoSelectionProps & { user: User; tariffs: Tariff[] }) {
  const [creditUser, setCreditUser] = useState(user.creditUser ?? "");
  const [creditPassword, setCreditPassword] = useState(
    user.creditPassword ?? ""
  );
  const [selectedTariff, setSelectedTariff] = useState<Tariff | null>(
    tariffs.find((t) => t.key === user.tariffKey) ?? null
  );
  const [cutOffTime, setCutOffTime] = useState(
    user.cutOffTime ? `${user.cutOffTime}` : ""
  );
  const [minimalCredit, setMinimalCredit] = useState(
    user.minimalCredit ? `${user.minimalCredit}` : ""
  );

  const [saveInfoToAccount, setSaveInfoToAccount] = useState(false);

  const updateUserDetails = useUpdateUserDetails();
  const { setToast } = useContext(ToastBarContext);

  const handleSubmit = (e: React.FormEvent) => {
    const data: Parameters<AutopurchaseInfoSelectionProps["handleSubmit"]>[1] =
      {
        creditPassword,
        creditUser,
        tariffKey: selectedTariff?.key ?? "",
        cutOffTime: cutOffTime !== "" ? Number(cutOffTime) : undefined,
        minimalCredit: minimalCredit !== "" ? Number(minimalCredit) : undefined,
      };

    if (saveInfoToAccount) {
      updateUserDetails.mutate(data, {
        onSuccess: () => {
          props.handleSubmit(e, data);
        },
        onError: (err) => {
          printError(setToast, err, "Failed to save details to account");
        },
      });
      return;
    }

    props.handleSubmit(e, data);
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

      <CreditNumber
        creditUser={creditUser}
        setCreditUser={setCreditUser}
        required
        error={creditUser === ""}
      />

      <CreditPassword
        creditPassword={creditPassword}
        setCreditPassword={setCreditPassword}
        error={creditPassword === ""}
        required
      />

      <SelectTariff
        selectedTariff={selectedTariff}
        setSelectedTariff={setSelectedTariff}
        tariffs={tariffs}
        required
        error={!selectedTariff}
      />

      <CutOffTime cutOffTime={cutOffTime} setCutOffTime={setCutOffTime} />

      <MinimalCredit
        minimalCredit={minimalCredit}
        setMinimalCredit={setMinimalCredit}
      />

      <FormControl fullWidth>
        <FormControlLabel
          control={
            <Checkbox
              checked={saveInfoToAccount}
              onChange={(e) => {
                setSaveInfoToAccount(e.target.checked);
              }}
              inputProps={{ "aria-label": "controlled" }}
            />
          }
          label="Save info to account?"
        />
      </FormControl>

      <SubmitButton
        isValid={isValid}
        tooltipTitle="First fill out all necessary info"
      >
        Start watching route
      </SubmitButton>
    </Box>
  );
}
