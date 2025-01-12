import { formatStation } from "@/utils/formatStation";
import { Station } from "@models/Station";
import { Autocomplete, Box, TextField, Typography } from "@mui/material";
import { DatePicker } from "@mui/x-date-pickers";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { Dayjs } from "dayjs";
import React, { Dispatch, SetStateAction } from "react";
import { useTranslation } from "react-i18next";
import SubmitButton from "./SubmitButton";

export default function StationAndDateSelection({
  stations,
  fromStation,
  setFromStation,
  toStation,
  setToStation,
  date,
  setDate,
  handleSubmit,
  isValid,
}: {
  stations: Station[];
  fromStation: Station | null;
  setFromStation: Dispatch<SetStateAction<Station | null>>;
  toStation: Station | null;
  setToStation: Dispatch<SetStateAction<Station | null>>;
  date: Dayjs | null;
  setDate: Dispatch<SetStateAction<Dayjs | null>>;
  handleSubmit: (event: React.FormEvent) => void;
  isValid: boolean;
}) {
  const { t } = useTranslation("default");

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
        <Typography variant="h5">{t("Select stations and date")}</Typography>
      </Box>
      <Autocomplete
        options={stations}
        getOptionLabel={(option) => formatStation(option)}
        value={fromStation}
        onChange={(_, newValue) => {
          setFromStation(newValue);
        }}
        renderInput={(params) => (
          <TextField {...params} label={t("From Station")} variant="outlined" />
        )}
        fullWidth
      />
      <Autocomplete
        options={stations}
        getOptionLabel={(option) => formatStation(option)}
        value={toStation}
        onChange={(_, newValue) => {
          setToStation(newValue);
        }}
        renderInput={(params) => (
          <TextField {...params} label={t("To Station")} variant="outlined" />
        )}
        fullWidth
      />
      <LocalizationProvider dateAdapter={AdapterDayjs}>
        <DatePicker
          label={t("Date")}
          value={date}
          onChange={(newValue) => {
            setDate(newValue);
          }}
          disablePast
          sx={{ width: "100%" }}
          timezone="UTC"
        />
      </LocalizationProvider>

      <SubmitButton
        isValid={isValid}
        tooltipTitle={t("First fill out all necessary info")}
      >
        {t("Search")}
      </SubmitButton>
    </Box>
  );
}
