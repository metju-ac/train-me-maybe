import config from "@/config";
import { User } from "@models/User";
import client from "@services/axiosClient";
import { getAuthToken, isUserLoggedIn, setAuthToken } from "@utils/authToken";

let users: (User & { token: string })[] = [];

const getCurrentUser = () => {
  if (!isUserLoggedIn()) {
    throw new Error("User is not logged in");
  }
  if (users.length === 0) {
    throw new Error("No users found");
  }
  const user = users.find((x) => x.token === getAuthToken());
  if (!user) {
    throw new Error("could not find user");
  }

  return user;
};

const userService = {
  getUserDetails: {
    key: "getUserDetails",
    fn: async () => {
      console.log(
        "getUserDetails: Getting details for the currently logged in user"
      );
      if (config.useMocks) {
        return getCurrentUser();
      }

      const response = await client.get<User>("/user");
      return response.data;
    },
  },

  updateUserDetails: {
    key: "updateUserDetails",
    fn: async (body: Omit<User, "email" | "debt">) => {
      console.log("Updating user details", body);

      if (config.useMocks) {
        users = [{ ...getCurrentUser(), ...body }];
        return body;
      }

      const response = await client.put<User>("/user", body);
      return response.data;
    },
  },

  registerUser: {
    key: "registerUser",
    fn: async (body: { email: string; password: string }) => {
      if (config.useMocks) {
        const token = "MockedToken" + (users.length + 1);
        setAuthToken(token);
        users.push({
          email: body.email,
          debt: 0,
          token: token,
        });
        return { token };
      }

      const response = await client.post<{ token: string }>("/user", body);

      setAuthToken(response.data.token);

      return response.data;
    },
  },

  loginUser: {
    key: "loginUser",
    fn: async (body: { email: string; password: string }) => {
      if (config.useMocks) {
        const user = users.find((user) => user.email === body.email);
        if (!user) {
          throw new Error("User not found");
        }
        setAuthToken(user.token);
        return { token: user.token };
      }

      const response = await client.post<{ token: string }>("/login", body);

      setAuthToken(response.data.token);

      return response.data;
    },
  },
} satisfies Record<
  string,
  { key: string; fn: (...args: any[]) => Promise<unknown> }
>;

export default userService;
