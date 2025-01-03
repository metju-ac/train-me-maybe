import { Route } from "@models/Route";
import { Station } from "@models/Station";
import {
  Box,
  Button,
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
import { SeatClass, Tariff } from "@services/staticDataService";
import useRoutes from "@utils/useRoutes";
import useSeatClasses from "@utils/useSeatClasses";
import useTariffs from "@utils/useTariffs";
import dayjs, { Dayjs } from "dayjs";
import utc from "dayjs/plugin/utc";
import React, { Dispatch, SetStateAction } from "react";

dayjs.extend(utc);

function formatRoute(route: Route) {
  return `${dayjs(route.departureTime).format("HH:mm:ss")} - ${dayjs(
    route.arrivalTime
  ).format("HH:mm:ss")} (${route.vehicleTypes}, price from ${
    route.creditPriceFrom
  })`;
}

interface RouteAndSeatAndTarriffSelectionProps {
  fromStation: Station;
  toStation: Station;
  date: Dayjs;
  handleSubmit: (event: React.FormEvent) => void;
  isValid: boolean;
  selectedRoute: Route | null;
  setSelectedRoute: Dispatch<SetStateAction<Route | null>>;
  selectedSeatClasses: SeatClass[];
  setSelectedSeatClasses: Dispatch<SetStateAction<SeatClass[]>>;
  selectedTariff: Tariff | null;
  setSelectedTariff: Dispatch<SetStateAction<Tariff | null>>;
  autopurchase: boolean;
  setAutopurchase: Dispatch<SetStateAction<boolean>>;
}

export default function RouteAndSeatAndTarriffSelection(
  props: RouteAndSeatAndTarriffSelectionProps
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

  const {
    data: tariffs,
    isLoading: isLoadingTariffs,
    isError: isErrorTariffs,
  } = useTariffs();

  if (isLoading || isLoadingSeatClasses || isLoadingTariffs) {
    return <CircularProgress />;
  }

  if (isError || isErrorSeatClasses || isErrorTariffs) {
    return <div>Error while loading routes</div>;
  }

  return (
    <RouteAndSeatAndTarriffSelectionForm
      {...props}
      routes={routes!}
      seatClasses={seatClasses!}
      tariffs={tariffs!}
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

function RouteAndSeatAndTarriffSelectionForm({
  date,
  handleSubmit,
  selectedRoute,
  setSelectedRoute,
  selectedSeatClasses,
  setSelectedSeatClasses,
  isValid,
  routes,
  seatClasses,
  tariffs,
  selectedTariff,
  setSelectedTariff,
  autopurchase,
  setAutopurchase,
}: RouteAndSeatAndTarriffSelectionProps & {
  routes: Route[];
  seatClasses: SeatClass[];
  tariffs: Tariff[];
}) {
  const handleChange: React.ChangeEventHandler<
    HTMLInputElement | HTMLTextAreaElement
  > = (event) => {
    const id = event.target.value;

    const route = routes?.find((route) => route.id === id) ?? null;

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
        <InputLabel id="select-label">Select a route</InputLabel>
        <Select
          labelId="select-label"
          name="option"
          value={selectedRoute?.id ?? ""}
          onChange={handleChange as any}
        >
          {routes.map((route) => (
            <MenuItem value={route.id} key={route.id}>
              {formatRoute(route)}
            </MenuItem>
          ))}
        </Select>
      </FormControl>

      <FormControl fullWidth>
        <InputLabel id="select-seat-classes">Select seat classes</InputLabel>
        <Select
          labelId="select-seat-classes"
          multiple
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
              onChange={(e) => setAutopurchase(e.target.checked)}
              inputProps={{ "aria-label": "controlled" }}
            />
          }
          label="Autopurchase ticket?"
        />
      </FormControl>

      {autopurchase && (
        <FormControl fullWidth>
          <InputLabel id="select-tariff">Select tariff</InputLabel>
          <Select
            labelId="select-tariff"
            value={selectedTariff?.key ?? ""}
            onChange={(e) =>
              setSelectedTariff(
                tariffs.find((t) => t.key === e.target.value) ?? null
              )
            }
          >
            {tariffs.map((tariff) => (
              <MenuItem key={tariff.key} value={tariff.key}>
                {tariff.key} ({tariff.value})
              </MenuItem>
            ))}
          </Select>
        </FormControl>
      )}

      <Tooltip title={isValid ? "" : "First fill out all necessary info"}>
        <span>
          <Button
            type="submit"
            variant="contained"
            color="primary"
            disabled={!isValid}
          >
            {autopurchase
              ? "Submit and review payment options"
              : "Start watching route"}
          </Button>
        </span>
      </Tooltip>
    </Box>
  );
}
