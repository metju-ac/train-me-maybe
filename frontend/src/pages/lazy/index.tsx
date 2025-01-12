import { lazy } from "react";

const LazyNewRoute = lazy(() => import("@pages/NewRoute"));
const LazyAbout = lazy(() => import("@pages/About"));
const LazyAccount = lazy(() => import("@pages/Account"));
const LazyHome = lazy(() => import("@pages/Home"));
const LazyLogin = lazy(() => import("@pages/Login"));
const LazyLogout = lazy(() => import("@pages/Logout"));
const LazyNotFound = lazy(() => import("@pages/NotFound"));
const LazyRegister = lazy(() => import("@pages/Register"));

export {
  LazyAbout,
  LazyAccount,
  LazyHome,
  LazyLogin,
  LazyLogout,
  LazyNewRoute,
  LazyNotFound,
  LazyRegister,
};
