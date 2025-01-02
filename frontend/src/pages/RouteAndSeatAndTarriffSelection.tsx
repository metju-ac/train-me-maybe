import { Route } from "@models/Route";
import { Station } from "@models/Station";
import {
  Box,
  Button,
  Checkbox,
  CircularProgress,
  FormControl,
  InputLabel,
  ListItemText,
  MenuItem,
  OutlinedInput,
  Select,
  SelectChangeEvent,
  TextField,
  Tooltip,
} from "@mui/material";
import useRoutes from "@utils/useRoutes";
import dayjs, { Dayjs } from "dayjs";
import utc from "dayjs/plugin/utc";
import React, { Dispatch, SetStateAction } from "react";

// Tell dayjs to use UTC timezone, since we are only concerned about dates, not times
dayjs.extend(utc);

function formatRoute(route: Route) {
  return `${dayjs(route.departureTime).format("HH:mm:ss")} - ${dayjs(
    route.arrivalTime
  ).format("HH:mm:ss")} (${route.vehicleTypes}, price from ${
    route.creditPriceFrom
  })`;
}

export default function RouteAndSeatAndTarriffSelection({
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
  const {
    data: routes,
    isLoading,
    isError,
  } = useRoutes({
    // convert to yyyy-mm-dd format
    date: date.format("YYYY-MM-DD"),
    fromStation: fromStation,
    toStation: toStation,
  });

  const handleChange: React.ChangeEventHandler<
    HTMLInputElement | HTMLTextAreaElement
  > = (event) => {
    const id = event.target.value;

    const route = routes?.find((route) => route.id === id) ?? null;

    setSelectedRoute(route);
  };

  const ITEM_HEIGHT = 48;
  const ITEM_PADDING_TOP = 8;
  const MenuProps = {
    PaperProps: {
      style: {
        maxHeight: ITEM_HEIGHT * 4.5 + ITEM_PADDING_TOP,
        width: 250,
      },
    },
  };

  const names = [
    "Oliver Hansen",
    "Van Henry",
    "April Tucker",
    "Ralph Hubbard",
    "Omar Alexander",
    "Carlos Abbott",
    "Miriam Wagner",
    "Bradley Wilkerson",
    "Virginia Andrews",
    "Kelly Snyder",
  ];
  const [personName, setPersonName] = React.useState<string[]>([]);
  const handleChangeTODO = (event: SelectChangeEvent<typeof personName>) => {
    const {
      target: { value },
    } = event;
    setPersonName(
      // On autofill we get a stringified value.
      typeof value === "string" ? value.split(",") : value
    );
  };

  if (isLoading) {
    return <CircularProgress />;
  }

  if (isError) {
    return <div>Error while loading routes</div>;
  }

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
      <TextField
        label="Selected date"
        value={date.format("YYYY-MM-DD")}
        fullWidth
        aria-readonly={true}
        disabled
      />
      <FormControl fullWidth>
        <InputLabel id="select-label">Select a route</InputLabel>
        <Select
          labelId="select-label"
          name="option"
          value={selectedRoute?.id ?? ""}
          onChange={handleChange as any}
        >
          {routes?.map((route) => (
            <MenuItem value={route.id} key={route.id}>
              {formatRoute(route)}
            </MenuItem>
          ))}
        </Select>
      </FormControl>

      <FormControl fullWidth>
        <InputLabel id="demo-multiple-checkbox-label">
          Select seat classes
        </InputLabel>
        <Select
          labelId="demo-multiple-checkbox-label"
          id="demo-multiple-checkbox"
          multiple
          value={personName}
          onChange={handleChangeTODO}
          input={<OutlinedInput label="Tag" />}
          renderValue={(selected) => selected.join(", ")}
          MenuProps={MenuProps}
        >
          {names.map((name) => (
            <MenuItem key={name} value={name}>
              <Checkbox checked={personName.includes(name)} />
              <ListItemText primary={name} />
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
    </Box>
  );
}
