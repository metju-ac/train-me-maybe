import axios from "axios";
import config from "@/config";
import { WatchedRoute } from "@/models/WatchedRoute";

const instance = axios.create({
  baseURL: config.baseUrl,
  timeout: 2000,
});

type Data = WatchedRoute;

function createData(
  id: number,
  tariffClass: string,
  fromStationId: number,
  routeId: number,
  toStationId: number,
  selectedSeatClasses: string[],
  userEmail: string
): Data {
  return {
    id,
    fromStationId,
    routeId,
    selectedSeatClasses,
    tariffClass,
    toStationId,
    userEmail,
  };
}

let rows = [
  createData(1, "Cupcake", 305, 3.7, 3.7, ["67"], "4.3"),
  createData(2, "Donut", 452, 25.0, 3.7, ["51"], "4.9"),
  createData(3, "Eclair", 262, 16.0, 3.7, ["24"], "6.0"),
];

const apiService = {
  getWatchedRoutes: {
    key: "getWatchedRoutes",
    fn: async () => {
      console.log("getWatchedRoutes: Getting all routes for current user");
      if (config.useMocks) {
        return rows;
      }
      try {
        const response = await instance.get<WatchedRoute[]>("/watchedRoute");
        return response.data;
      } catch (error) {
        console.error(error);
        return [];
      }
    },
  },
  deleteWatchedRoute: {
    key: "deleteWatchedRoute",
    fn: async (ids: readonly number[]) => {
      if (config.useMocks) {
        rows = rows.filter((row) => !ids.includes(row.id));
        return;
      }
      try {
        const promises = ids.map((id) =>
          instance.delete(`/watchedRoute/${id}`)
        );

        await Promise.all(promises);
      } catch (error) {
        console.error(error);
      }
    },
  },
} satisfies Record<
  string,
  { key: string; fn: (...args: any[]) => Promise<unknown> }
>;

export default apiService;
