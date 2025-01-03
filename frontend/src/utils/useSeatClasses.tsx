import staticDataService from "@services/staticDataService";
import { useQuery } from "@tanstack/react-query";

export default function useSeatClasses() {
  return useQuery({
    queryKey: [staticDataService.getSeatClasses.key],
    queryFn: staticDataService.getSeatClasses.fn,
  });
}
