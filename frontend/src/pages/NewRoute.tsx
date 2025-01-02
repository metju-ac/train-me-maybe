import config from "@/config";
import { Route } from "@models/Route";
import { Station } from "@models/Station";
import {
  Autocomplete,
  Box,
  Button,
  CircularProgress,
  FormControl,
  InputLabel,
  MenuItem,
  Select,
  TextField,
  Tooltip,
} from "@mui/material";
import { DatePicker } from "@mui/x-date-pickers";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import useRoutes from "@utils/useRoutes";
import useStations from "@utils/useStations";
import dayjs, { Dayjs } from "dayjs";
import utc from "dayjs/plugin/utc";
import React, { Dispatch, SetStateAction, useState } from "react";

// Tell dayjs to use UTC timezone, since we are only concerned about dates, not times
dayjs.extend(utc);

function StationAndDateSelection({
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

function RouteSelection({
  fromStation,
  toStation,
  date,
  handleSubmit,
  selectedRoute,
  setSelectedRoute,
  isValid,
}: {
  fromStation: Station;
  toStation: Station;
  date: Dayjs;
  handleSubmit: (event: React.FormEvent) => void;
  isValid: boolean;
  selectedRoute: Route | null;
  setSelectedRoute: Dispatch<SetStateAction<Route | null>>;
}) {
  console.log("date", date, "string", date.toISOString().substring(0, 10));

  const {
    data: routes,
    isLoading,
    isError,
  } = useRoutes({
    // convert to yyyy-mm-dd format
    date: date.toISOString().substring(0, 10),
    fromStation: fromStation,
    toStation: toStation,
  });

  if (isLoading) {
    return <CircularProgress />;
  }

  if (isError) {
    return <div>Error while loading routes</div>;
  }

  const handleChange = (event: any) => {
    console.log("handle change", event);
    // const { name, value } = event.target;
    // setFormValues({
    //   ...formValues,
    //   [name]: value
    // });
  };

  return (
    <form onSubmit={handleSubmit}>
      <TextField
        label="Name"
        name="name"
        value={"yyy" + (selectedRoute?.departureTime ?? "")}
        onChange={handleChange}
        fullWidth
        margin="normal"
      />
      <FormControl fullWidth margin="normal">
        <InputLabel id="select-label">Options</InputLabel>
        <Select
          labelId="select-label"
          name="option"
          value={selectedRoute}
          onChange={handleChange}
        >
          {routes?.map((route) => (
            <MenuItem value={route.id}>
              xxx route {route.departureTime}
            </MenuItem>
          ))}
        </Select>
      </FormControl>
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
    </form>
  );
}

export default function NewRoute() {
  // First, there is a form with selection of stations and a date of departure
  const { data: stations, isLoading, isError } = useStations();

  if (isLoading) {
    return <CircularProgress />;
  }

  if (isError) {
    return <div>Error while fetching data</div>;
  }

  return <NewRouteForm stations={stations as Station[]} />;
}

function NewRouteForm({ stations }: { stations: Station[] }) {
  // we will mimick the regiojet mobile app search
  const [fromStation, setFromStation] = useState<Station | null>(() => {
    if (config.useMocks) {
      return stations.find((station) => station.stationID === 10204002) ?? null;
    }
    return null;
  });
  const [toStation, setToStation] = useState<Station | null>(() => {
    if (config.useMocks) {
      return stations.find((station) => station.stationID === 10204003) ?? null;
    }
    return null;
  });
  const [date, setDate] = useState<Dayjs | null>(() => {
    if (config.useMocks) {
      return dayjs.utc("2025-01-21");
    }
    return null;
  });
  const [selectedRoute, setSelectedRoute] = useState<Route | null>(null);

  const [phase, setPhase] = useState<"stations" | "route" | "ticket">(
    "stations"
  );

  const handleSubmitStationsPhase = (event: React.FormEvent) => {
    event.preventDefault();
    // Handle form submission logic here
    console.log("From Station:", fromStation);
    console.log("To Station:", toStation);
    console.log("Date:", date);

    setPhase("route");
  };

  const handleSubmitRoutePhase = (event: React.FormEvent) => {
    event.preventDefault();
    // Handle form submission logic here
    console.log("From Station:", fromStation);
    console.log("To Station:", toStation);
    console.log("Date:", date);
    console.log("Selected route:", "TODO");

    setPhase("ticket");
  };

  const renderContent = () => {
    switch (phase) {
      case "stations":
        return (
          <StationAndDateSelection
            date={date}
            fromStation={fromStation}
            handleSubmit={handleSubmitStationsPhase}
            setDate={setDate}
            setFromStation={setFromStation}
            setToStation={setToStation}
            stations={stations as Station[]}
            toStation={toStation}
            isValid={
              toStation !== null && fromStation !== null && date !== null
            }
          />
        );
      case "route":
        return (
          <RouteSelection
            date={date!}
            fromStation={fromStation!}
            toStation={toStation!}
            handleSubmit={handleSubmitRoutePhase}
            isValid={selectedRoute !== null}
            selectedRoute={selectedRoute}
            setSelectedRoute={setSelectedRoute}
          />
        );
      case "ticket":
        return <div>world</div>;
      default:
        return <div>Unknown phase</div>;
    }
  };

  return renderContent();
}
