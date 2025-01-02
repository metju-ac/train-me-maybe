import staticDataService from "@services/staticDataService";
import { useQuery } from "@tanstack/react-query";

export default function useStations() {
  return useQuery({
    queryKey: [staticDataService.getStations.key],
    queryFn: staticDataService.getStations.fn,
  });
}
