import config from "@/config";
import { Route } from "@models/Route";
import { Station } from "@models/Station";
import { formatStation } from "@/utils/formatStation";
import client from "@services/axiosClient";
import routes from "./data/routes.json";

const routeService = {
  getRoutes: {
    key: "getRoutes",
    fn: async (params: {
      fromStation: Station;
      toStation: Station;
      date: string;
    }): Promise<Route[]> => {
      console.log(
        "getRoutes: getting routes from station",
         formatStation(params.fromStation),
        "to station",
         formatStation(params.toStation),
        "on date",
        params.date
      );

      if (config.useMocks) {
        // "departureTime": "2025-01-02T22:00:00.000+01:00" -> get only the first 10 characters
        const matchingRoutes = [];

        for (const route of routes.routes) {
          if (
            route.departureStationId === params.fromStation.stationId &&
            route.arrivalStationId === params.toStation.stationId &&
            params.date === route.departureTime.substring(0, 10)
          ) {
            matchingRoutes.push(route);
          }
        }

        if (matchingRoutes.length > 0) {
          return matchingRoutes;
        }

        throw new Error("No routes found");
      }

      const response = await client.get<Route[]>("/auth/route", {
        params: {
          fromStationId: params.fromStation.stationId,
          toStationId: params.toStation.stationId,
          date: params.date,
        },
      });
      return response.data;
    },
  },
} satisfies Record<
  string,
  { key: string; fn: (...args: any[]) => Promise<unknown> }
>;

export default routeService;
