import { useSyncExternalStore } from "react";
import {
  AUTH_TOKEN_KEY,
  isUserLoggedIn,
  subscribeToAuthTokenChanges,
} from "./authToken";

function getSnapshot() {
  console.log("getSnapshot called");
  return isUserLoggedIn();
}

/**
 * Subscribes to changes in auth token, both in the current browsing context (this tab)
 * and in other tabs.
 */
function subscribe(callback: () => void) {
  const browserCallback = (e: StorageEvent) => {
    if (e.key !== AUTH_TOKEN_KEY) {
      console.log("Ignoring storage event for key", e.key);
      return;
    }

    console.log("Accepting auth change from a different tab: key", e.key);
    callback();
  };

  const unsubscribe = subscribeToAuthTokenChanges(() => {
    console.log("Accepting auth change from the current tab");
    callback();
  });

  window.addEventListener("storage", browserCallback);
  return () => {
    unsubscribe();
    window.removeEventListener("storage", browserCallback);
  };
}

/**
 * A hook which returns true if the user is logged in, false otherwise.
 */
export default function useIsLoggedIn() {
  return useSyncExternalStore(subscribe, getSnapshot);
}
