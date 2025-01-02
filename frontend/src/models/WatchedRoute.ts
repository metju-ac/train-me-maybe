export interface WatchedRoute {
  id: number;
  userEmail: string;
  fromStationId: number;
  toStationId: number;
  routeId: number;
  tariffClass: string;
  selectedSeatClasses: string[];
}
