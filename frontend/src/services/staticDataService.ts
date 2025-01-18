import config from "@/config";
import { AvailableLanguage } from "@/i18n";
import { Station } from "@models/Station";
import client from "./axiosClient";
import seatClasses from "./data/seatClasses.json";
import countries from "./data/stations.json";
import tariffs from "./data/tariffs.json";

export type SeatClass = (typeof seatClasses)[0];
export type Tariff = (typeof tariffs)[0];

function transformStationsFromRaw(): Station[] {
  const result: Station[] = [];

  for (const country of countries) {
    for (const city of country.cities) {
      for (const station of city.stations) {
        let isTrainStation = false;
        let isBusStation = false;
        for (const stationType of station.stationsTypes) {
          if (stationType === "TRAIN_STATION") {
            isTrainStation = true;
          }
          if (stationType === "BUS_STATION") {
            isBusStation = true;
          }
        }

        if (!isTrainStation && !isBusStation) {
          console.warn(
            "Station is neither a train station nor a bus station",
            "stationName",
            station.name,
            "stationID",
            station.id
          );
        }

        result.push({
          country: country.country,
          city: city.name,
          stationID: station.id,
          stationName: station.fullname,
          isTrainStation,
          isBusStation,
        });
      }
    }
  }

  return result;
}

const staticDataService = {
  getSeatClasses: {
    key: "getSeatClasses",
    fn: async (params: {
      fromStationId: number;
      toStationId: number;
      routeId: string;
    }): Promise<SeatClass[]> => {
      console.log("getSeatClasses: Getting all seat classes for route", params);
      const response = await client.get<{
        seatClasses: SeatClass["key"][];
      }>("/auth/seatClass", {
        params,
      });
      const availableClasses = response.data.seatClasses;

      return seatClasses.filter((seatClass) =>
        availableClasses.includes(seatClass.key)
      );
    },
  },

  getTariffs: {
    key: "getTariffs",
    fn: async (): Promise<Tariff[]> => {
      console.log("getTariffs: Getting all tariffs");
      return tariffs;
    },
  },

  getStations: {
    key: "getStations",
    fn: async (lang: AvailableLanguage = "en"): Promise<Station[]> => {
      console.log("getStations: Getting all stations");
      if (config.useMocks) {
        return transformStationsFromRaw();
      }
      try {
        const response = await client.get<Station[]>("/auth/station", {
          params: {
            lang,
          },
        });
        return response.data;
      } catch (error) {
        console.error(error);
        return [];
      }
    },
  },
} satisfies Record<
  string,
  { key: string; fn: (...args: any[]) => Promise<unknown> }
>;

export default staticDataService;
