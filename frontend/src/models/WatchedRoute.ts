export interface WatchedRoute {
  id: number;
  userEmail: string;
  fromStationId: number;
  toStationId: number;
  routeId: string;
  tariffClass: string;
  selectedSeatClasses: string[];
}
