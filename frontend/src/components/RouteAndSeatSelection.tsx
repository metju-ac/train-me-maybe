import { Route } from "@models/Route";
import { Station } from "@models/Station";
import {
  Box,
  Checkbox,
  CircularProgress,
  FormControl,
  FormControlLabel,
  InputLabel,
  ListItemText,
  MenuItem,
  OutlinedInput,
  Select,
  SelectChangeEvent,
  TextField,
  Tooltip,
  Typography,
} from "@mui/material";
import { SeatClass } from "@services/staticDataService";
import useRoutes from "@utils/useRoutes";
import useSeatClasses from "@utils/useSeatClasses";
import dayjs, { Dayjs } from "dayjs";
import utc from "dayjs/plugin/utc";
import React, { Dispatch, ReactNode, SetStateAction } from "react";
import SubmitButton from "./SubmitButton";

dayjs.extend(utc);

function formatRoute(route: Route) {
  return `${dayjs(route.departureTime).format("HH:mm:ss")} - ${dayjs(
    route.arrivalTime
  ).format("HH:mm:ss")} (${route.vehicleTypes.join(", ")}, price from ${
    route.creditPriceFrom
  })`;
}

interface RouteAndSeatSelectionProps {
  fromStation: Station;
  toStation: Station;
  date: Dayjs;
  handleSubmit: (event: React.FormEvent) => void;
  isValid: boolean;
  selectedRoute: Route | null;
  setSelectedRoute: Dispatch<SetStateAction<Route | null>>;
  selectedSeatClasses: SeatClass[];
  setSelectedSeatClasses: Dispatch<SetStateAction<SeatClass[]>>;
  autopurchase: boolean;
  setAutopurchase: Dispatch<SetStateAction<boolean>>;
}

export default function RouteAndSeatSelection(
  props: RouteAndSeatSelectionProps
) {
  const {
    data: routes,
    isLoading,
    isError,
  } = useRoutes({
    // convert to yyyy-mm-dd format
    date: props.date.format("YYYY-MM-DD"),
    fromStation: props.fromStation,
    toStation: props.toStation,
  });

  const {
    data: seatClasses,
    isLoading: isLoadingSeatClasses,
    isError: isErrorSeatClasses,
  } = useSeatClasses();

  if (isLoading || isLoadingSeatClasses) {
    return <CircularProgress />;
  }

  if (isError || isErrorSeatClasses) {
    return <div>Error while loading routes</div>;
  }

  return (
    <RouteAndSeatSelectionForm
      {...props}
      routes={routes!}
      seatClasses={seatClasses!}
    />
  );
}

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

function renderSeatClass(key: string, seatClasses: SeatClass[]) {
  const seatClass = seatClasses.find((s) => s.key === key);
  return seatClass?.title ?? key;
}

function RouteAndSeatSelectionForm({
  date,
  handleSubmit,
  selectedRoute,
  setSelectedRoute,
  selectedSeatClasses,
  setSelectedSeatClasses,
  isValid,
  routes,
  seatClasses,
  autopurchase,
  setAutopurchase,
}: RouteAndSeatSelectionProps & {
  routes: Route[];
  seatClasses: SeatClass[];
}) {
  const handleChange: (
    event: SelectChangeEvent<string>,
    child: ReactNode
  ) => void = (event) => {
    const id = event.target.value;

    const route = routes.find((route) => route.id === id) ?? null;

    setSelectedRoute(route);
  };

  const handleChangeSeatClasses = (
    event: SelectChangeEvent<SeatClass["key"][]>
  ) => {
    const {
      target: { value },
    } = event;

    if (typeof value === "string") {
      const keys = value.split(",");
      setSelectedSeatClasses(seatClasses.filter((s) => keys.includes(s.key)));
      return;
    }

    const selected = seatClasses.filter((s) => value.includes(s.key));
    setSelectedSeatClasses(selected);
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
      <Box>
        <Typography variant="h5">
          Select route, seat classes and tariff
        </Typography>
      </Box>
      <TextField
        label="Selected date"
        value={date.format("YYYY-MM-DD")}
        fullWidth
        aria-readonly={true}
        disabled
      />
      <FormControl fullWidth>
        <InputLabel id="select-label">Route</InputLabel>
        <Select
          labelId="select-label"
          name="option"
          label="Route"
          value={selectedRoute?.id ?? ""}
          onChange={handleChange}
        >
          {routes.map((route) => (
            <MenuItem value={route.id} key={route.id}>
              {formatRoute(route)}
            </MenuItem>
          ))}
        </Select>
      </FormControl>

      <FormControl fullWidth>
        <InputLabel id="select-seat-classes">Seat classes</InputLabel>
        <Select
          labelId="select-seat-classes"
          multiple
          label="Seat classes"
          value={selectedSeatClasses.map((s) => s.key)}
          onChange={handleChangeSeatClasses}
          input={<OutlinedInput label="Tag" />}
          renderValue={(selected) =>
            selected.map((s) => renderSeatClass(s, seatClasses)).join(", ")
          }
          MenuProps={MenuProps}
        >
          {seatClasses.map((seatClass) => (
            <MenuItem key={seatClass.key} value={seatClass.key}>
              <Checkbox checked={selectedSeatClasses.includes(seatClass)} />
              <ListItemText primary={seatClass.title} />
            </MenuItem>
          ))}
        </Select>
      </FormControl>

      <FormControl fullWidth>
        <FormControlLabel
          control={
            <Checkbox
              checked={autopurchase}
              onChange={(e) => {
                setAutopurchase(e.target.checked);
              }}
              inputProps={{ "aria-label": "controlled" }}
            />
          }
          label={
            <Tooltip
              arrow
              title="If checked, we will try to purchase the ticket for you. You will need to provide Regiojet credentials, and you need to have sufficient credit."
            >
              <Typography>Autopurchase ticket?</Typography>
            </Tooltip>
          }
        />
      </FormControl>

      <SubmitButton
        isValid={isValid}
        tooltipTitle="First fill out all necessary info"
      >
        {autopurchase ? "Review payment options" : "Start watching route"}
      </SubmitButton>
    </Box>
  );
}
