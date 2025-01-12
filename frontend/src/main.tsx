import AuthenticatedOnly from "@components/AuthenticatedOnly";
import ToastBarProvider from "@components/ToastBarProvider";
import {
  LazyAbout,
  LazyAccount,
  LazyHome,
  LazyLogin,
  LazyLogout,
  LazyNewRoute,
  LazyNotFound,
  LazyRegister,
} from "@pages/lazy";
import Root from "@pages/Root.tsx";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { BrowserRouter, Route, Routes } from "react-router";
import { initI18N } from "./i18n";
import "./index.css";

// Create a client
const queryClient = new QueryClient();

const home = <AuthenticatedOnly element={<LazyHome />} />;
const account = <AuthenticatedOnly element={<LazyAccount />} />;
const logout = <AuthenticatedOnly element={<LazyLogout />} />;
const newRoute = <AuthenticatedOnly element={<LazyNewRoute />} />;

initI18N()
  .then(() =>
    createRoot(document.getElementById("root")!).render(
      <StrictMode>
        <QueryClientProvider client={queryClient}>
          <ToastBarProvider>
            <BrowserRouter>
              <Routes>
                <Route path="/" element={<Root />}>
                  <Route index element={home} />
                  <Route path="routes/new" element={newRoute} />
                  <Route path="about" element={<LazyAbout />} />
                  <Route path="account" element={account} />
                  <Route path="login" element={<LazyLogin />} />
                  <Route path="register" element={<LazyRegister />} />
                  <Route path="logout" element={logout} />
                  <Route path="*" element={<LazyNotFound />} />
                </Route>
              </Routes>
            </BrowserRouter>
          </ToastBarProvider>
        </QueryClientProvider>
      </StrictMode>
    )
  )
  .catch(console.error);
