import config from "@/config";
import { Route } from "@models/Route";
import { Station } from "@models/Station";
import { CircularProgress } from "@mui/material";
import useStations from "@utils/useStations";
import dayjs, { Dayjs } from "dayjs";
import React, { useState } from "react";
import RouteAndSeatAndTarriffSelection from "./RouteAndSeatAndTarriffSelection";
import StationAndDateSelection from "./StationAndDateSelection";

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

  const [phase, setPhase] = useState<
    "stations-date" | "route-seatClasses-tariff" | "autopurchase"
  >("stations-date");

  const handleSubmitStationsPhase = (event: React.FormEvent) => {
    event.preventDefault();
    // Handle form submission logic here
    console.log("From Station:", fromStation);
    console.log("To Station:", toStation);
    console.log("Date:", date);

    setPhase("route-seatClasses-tariff");
  };

  const handleSubmitRoutePhase = (event: React.FormEvent) => {
    event.preventDefault();
    // Handle form submission logic here
    console.log("From Station:", fromStation);
    console.log("To Station:", toStation);
    console.log("Date:", date);
    console.log("Selected route:", "TODO");

    setPhase("autopurchase");
  };

  const renderContent = () => {
    switch (phase) {
      case "stations-date":
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
      case "route-seatClasses-tariff":
        return (
          <RouteAndSeatAndTarriffSelection
            date={date!}
            fromStation={fromStation!}
            toStation={toStation!}
            handleSubmit={handleSubmitRoutePhase}
            isValid={selectedRoute !== null} // TODO
            selectedRoute={selectedRoute}
            setSelectedRoute={setSelectedRoute}
          />
        );
      case "autopurchase":
        return <div>world</div>;
      default:
        return <div>Unknown phase</div>;
    }
  };

  return renderContent();
}
