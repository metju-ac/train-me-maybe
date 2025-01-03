import { User } from "@models/User";
import { Box, Typography } from "@mui/material";
import { Tariff } from "@services/staticDataService";
import useUpdateUserDetails from "@utils/useUpdateUserDetails";
import React, { useState } from "react";
import CreditNumber from "./CreditNumber";
import CreditPassword from "./CreditPassword";
import CutOffTime from "./CutOffTime";
import MinimalCredit from "./MinimalCredit";
import SelectTariff from "./SelectTariff";
import SubmitButton from "./SubmitButton";

export default function AccountForm({
  user,
  tariffs,
}: {
  user: User;
  tariffs: Tariff[];
}) {
  const [creditUser, setCreditUser] = useState(user.creditUser || "");
  const [creditPassword, setCreditPassword] = useState(
    user.creditPassword || ""
  );
  const [selectedTariff, setSelectedTariff] = useState<Tariff | null>(
    tariffs.find((t) => t.key === user.tariffKey) ?? null
  );
  const [cutOffTime, setCutOffTime] = useState(`${user.cutOffTime}` || "");
  const [minimalCredit, setMinimalCredit] = useState(
    `${user.minimalCredit}` || ""
  );

  const hasChanged =
    creditUser !== user.creditUser ||
    creditPassword !== user.creditPassword ||
    cutOffTime !== `${user.cutOffTime}` ||
    selectedTariff?.key !== user.tariffKey ||
    minimalCredit !== `${user.minimalCredit}`;

  const mutation = useUpdateUserDetails();

  const handleSubmit = (event: React.FormEvent) => {
    event.preventDefault();
    mutation.mutate({
      creditUser,
      creditPassword,
      cutOffTime: cutOffTime !== "" ? Number(cutOffTime) : undefined,
      tariffKey: selectedTariff?.key ?? undefined,
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
        gap: 2,
      }}
    >
      <Typography variant="h4" component="h1" gutterBottom>
        Account Details
      </Typography>
      <Typography variant="h6" component="h2" gutterBottom>
        Email: {user.email}
      </Typography>

      <CreditNumber creditUser={creditUser} setCreditUser={setCreditUser} />

      <CreditPassword
        creditPassword={creditPassword}
        setCreditPassword={setCreditPassword}
      />
      <SelectTariff
        selectedTariff={selectedTariff}
        setSelectedTariff={setSelectedTariff}
        tariffs={tariffs}
      />

      <CutOffTime cutOffTime={cutOffTime} setCutOffTime={setCutOffTime} />

      <MinimalCredit
        minimalCredit={minimalCredit}
        setMinimalCredit={setMinimalCredit}
      />

      <SubmitButton isValid={hasChanged} tooltipTitle="No changes were made">
        Update Details
      </SubmitButton>
    </Box>
  );
}
