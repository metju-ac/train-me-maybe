import { removeAuthToken } from "@utils/authToken";
import { useEffect } from "react";

export default function Logout() {
  useEffect(() => {
    removeAuthToken();
  }, []);

  return (
    <>
      <h1>Logout</h1>
      <p>You have been logged out.</p>
    </>
  );
}
