import config from "@/config";
import { WatchedRoute } from "@/models/WatchedRoute";
import client from "@services/axiosClient";

function createData(
  id: number,
  tariffClass: string,
  fromStationId: number,
  routeId: number,
  toStationId: number,
  selectedSeatClasses: string[],
  userEmail: string
): WatchedRoute {
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
  createData(1, "Cupcake", 4961583004, 3.7, 4987881000, ["67"], "4.3"),
  createData(2, "Donut", 5095524063, 25.0, 4987881000, ["51"], "4.9"),
  createData(3, "Eclair", 4961583004, 16.0, 5095524063, ["24"], "6.0"),
];

const watchedRouteService = {
  getWatchedRoutes: {
    key: "getWatchedRoutes",
    fn: async () => {
      console.log("getWatchedRoutes: Getting all routes for current user");
      if (config.useMocks) {
        return rows;
      }
      try {
        const response = await client.get<WatchedRoute[]>("/watchedRoute");
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
        const promises = ids.map((id) => client.delete(`/watchedRoute/${id}`));

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

export default watchedRouteService;