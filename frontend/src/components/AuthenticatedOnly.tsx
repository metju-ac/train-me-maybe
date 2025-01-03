import useIsLoggedIn from "@utils/useIsLoggedIn";
import React from "react";

export default function AuthenticatedOnly(props: { element: React.ReactNode }) {
  const isLoggedIn = useIsLoggedIn();

  if (!isLoggedIn) {
    return (
      <>
        <h1>You are not logged in</h1>
        <p>To use the app you need to register.</p>
      </>
    );
  }

  return props.element;
}
