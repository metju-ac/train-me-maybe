const listeners: (() => void)[] = [];
export const AUTH_TOKEN_KEY = "authToken";

/**
 * Subscribe to changes in auth token.
 *
 * Returns a function that can be called to unsubscribe.
 */
export const subscribeToAuthTokenChanges = (callback: () => void) => {
  listeners.push(callback);
  return () => {
    const index = listeners.indexOf(callback);
    listeners.splice(index, 1);
  };
};

/**
 * Gets auth token from local storage.
 */
export const getAuthToken = () => {
  return localStorage.getItem(AUTH_TOKEN_KEY);
};

export const setAuthToken = (token: string) => {
  localStorage.setItem(AUTH_TOKEN_KEY, token);
  listeners.forEach((listener) => listener());
};

export const removeAuthToken = () => {
  localStorage.removeItem(AUTH_TOKEN_KEY);
  listeners.forEach((listener) => listener());
};

export const isUserLoggedIn = () => {
  return !!getAuthToken();
};
