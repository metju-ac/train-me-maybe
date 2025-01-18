export interface WatchedRoute {
  id: number;
  userEmail: string;
  fromStationId: number;
  toStationId: number;
  routeId: string;
  selectedSeatClasses: string[];
  autoPurchase: boolean;
  tariffClass: string;
  cutOffTime: number | null;
  minimalCredit: number | null;
  departureDateTime: string;
}
