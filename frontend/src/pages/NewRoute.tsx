import config from "@/config";
import AutopurchaseInfoSelection, {
  AutopurchaseInfoSelectionProps,
} from "@components/AutopurchaseInfoSelection";
import { Route } from "@models/Route";
import { Station } from "@models/Station";
import { CircularProgress } from "@mui/material";
import { SeatClass } from "@services/staticDataService";
import useCreateWatchedRoute from "@utils/useCreateWatchedRoute";
import useStations from "@utils/useStations";
import dayjs, { Dayjs } from "dayjs";
import React, { useState } from "react";
import RouteAndSeatSelection from "../components/RouteAndSeatSelection";
import StationAndDateSelection from "../components/StationAndDateSelection";

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
  const [selectedSeatClasses, setSelectedSeatClasses] = useState<SeatClass[]>(
    []
  );
  const [autopurchase, setAutopurchase] = useState(false);

  const [phase, setPhase] = useState<
    "stations-date" | "route-seatClasses-tariff" | "autopurchase"
  >("stations-date");

  const mutation = useCreateWatchedRoute();

  const handleSubmitStationsPhase = (event: React.FormEvent) => {
    event.preventDefault();
    setPhase("route-seatClasses-tariff");
  };

  const handleSubmitAutopurchasePhase: AutopurchaseInfoSelectionProps["handleSubmit"] =
    (event, data) => {
      event.preventDefault();

      mutation.mutate({
        autoPurchase: true,
        fromStationId: fromStation!.stationID,
        toStationId: toStation!.stationID,
        routeId: selectedRoute!.id,
        selectedSeatClasses: selectedSeatClasses.map(
          (seatClass) => seatClass.key
        ),
        tariffClass: data.tariffKey,
        creditUser: data.creditUser,
        creditPassword: data.creditPassword,
        cutOffTime: data.cutOffTime,
        minimalCredit: data.minimalCredit,
      });
    };

  const handleSubmitRoutePhase = (event: React.FormEvent) => {
    event.preventDefault();

    if (autopurchase) {
      setPhase("autopurchase");
      return;
    }

    mutation.mutate({
      autoPurchase: false,
      fromStationId: fromStation!.stationID,
      toStationId: toStation!.stationID,
      routeId: selectedRoute!.id,
      selectedSeatClasses: selectedSeatClasses.map(
        (seatClass) => seatClass.key
      ),
    });
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
            stations={stations}
            toStation={toStation}
            isValid={
              toStation !== null && fromStation !== null && date !== null
            }
          />
        );
      case "route-seatClasses-tariff":
        return (
          <RouteAndSeatSelection
            date={date!}
            fromStation={fromStation!}
            toStation={toStation!}
            handleSubmit={handleSubmitRoutePhase}
            isValid={selectedRoute !== null && selectedSeatClasses.length > 0}
            selectedRoute={selectedRoute}
            setSelectedRoute={setSelectedRoute}
            selectedSeatClasses={selectedSeatClasses}
            setSelectedSeatClasses={setSelectedSeatClasses}
            autopurchase={autopurchase}
            setAutopurchase={setAutopurchase}
          />
        );
      case "autopurchase":
        return (
          <AutopurchaseInfoSelection
            handleSubmit={handleSubmitAutopurchasePhase}
          />
        );
      default:
        return <div>Unknown phase</div>;
    }
  };

  return renderContent();
}
