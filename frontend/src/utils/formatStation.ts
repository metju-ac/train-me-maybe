import { Station } from "@/models/Station";

export function formatStation(station: Station): string {
  return `${station.city} - ${station.stationName}`;
}
