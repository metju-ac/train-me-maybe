import config from "@/config";
import { User } from "@models/User";
import client from "@services/axiosClient";

let users: User[] = [];

const userService = {
  getUserDetails: {
    key: "getUserDetails",
    fn: async () => {
      console.log(
        "getUserDetails: Getting details for the currently logged in user"
      );
      if (config.useMocks) {
        if (users.length === 0) {
          throw new Error("No users found");
        }
        return users[0];
      }
      try {
        const response = await client.get<User>("/user");
        return response.data;
      } catch (error) {
        console.error(error);
        return [];
      }
    },
  },

  registerUser: {
    key: "registerUser",
    fn: async (body: { email: string; password: string }) => {
      if (config.useMocks) {
        users.push({ email: body.email, debt: 0 });
        return { token: "mockedToken" };
      }

      const token = await client.post<{ token: string }>("/user", body);
      return token.data;
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
        return { token: "mockedToken" };
      }

      const token = await client.post<{ token: string }>("/login", body);
      return { success: true, data: token.data };
    },
  },
} satisfies Record<
  string,
  { key: string; fn: (...args: any[]) => Promise<unknown> }
>;

export default userService;
