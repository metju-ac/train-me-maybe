import config from "@/config";
import { User } from "@models/User";
import client from "@services/axiosClient";

let users: User[] = [{ email: "a@a.com", debt: 0 }];

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

      const response = await client.get<User>("/user");
      return response.data;
    },
  },

  updateUserDetails: {
    key: "updateUserDetails",
    fn: async (body: Omit<User, "email" | "debt">) => {
      console.log("Updating user details", body);

      if (config.useMocks) {
        users = [{ ...users[0], ...body }];
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
