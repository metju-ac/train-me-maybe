export interface Route {
  id: string;
  departureStationId: number;
  departureTime: string;
  arrivalStationId: number;
  arrivalTime: string;
  vehicleTypes: string[];
  freeSeatsCount: number;
  priceFrom: number;
  priceTo: number;
  creditPriceFrom: number;
  creditPriceTo: number;
  notices: boolean;
  nationalTrip: boolean;
  bookable: boolean;
  travelTime: string;
}
