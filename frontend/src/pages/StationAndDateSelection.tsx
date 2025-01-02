import { Station } from "@models/Station";
import { Autocomplete, Box, Button, TextField, Tooltip } from "@mui/material";
import { DatePicker } from "@mui/x-date-pickers";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { Dayjs } from "dayjs";
import React, { Dispatch, SetStateAction } from "react";

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
      <Autocomplete
        options={stations!}
        getOptionLabel={(option) => option.stationName}
        value={fromStation}
        onChange={(_, newValue) => setFromStation(newValue)}
        renderInput={(params) => (
          <TextField {...params} label="From Station" variant="outlined" />
        )}
        fullWidth
      />
      <Autocomplete
        options={stations!}
        getOptionLabel={(option) => option.stationName}
        value={toStation}
        onChange={(_, newValue) => setToStation(newValue)}
        renderInput={(params) => (
          <TextField {...params} label="To Station" variant="outlined" />
        )}
        fullWidth
      />
      <LocalizationProvider dateAdapter={AdapterDayjs}>
        <DatePicker
          label="Date"
          value={date}
          onChange={(newValue) => setDate(newValue)}
          disablePast
          sx={{ width: "100%" }}
          timezone="UTC"
        />
      </LocalizationProvider>
      <Tooltip title={isValid ? "" : "First fill out all necessary info"}>
        <span>
          <Button
            type="submit"
            variant="contained"
            color="primary"
            disabled={!isValid}
          >
            Search
          </Button>
        </span>
      </Tooltip>
    </Box>
  );
}
