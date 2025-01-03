import userService from "@services/userService";
import { useQuery } from "@tanstack/react-query";

export default function useUserDetails() {
  return useQuery({
    queryKey: [userService.getUserDetails.key],
    queryFn: userService.getUserDetails.fn,
    retry: 2,
  });
}
