import ToastBarProvider from "@components/ToastBarProvider";
import About from "@pages/About.tsx";
import Account from "@pages/Account.tsx";
import Home from "@pages/Home.tsx";
import Login from "@pages/Login";
import Logout from "@pages/Logout";
import NewRoute from "@pages/NewRoute";
import NotFound from "@pages/NotFound";
import Register from "@pages/Register";
import Root from "@pages/Root.tsx";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { BrowserRouter, Route, Routes } from "react-router";
import "./index.css";

// Create a client
const queryClient = new QueryClient();

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <QueryClientProvider client={queryClient}>
      <ToastBarProvider>
        <BrowserRouter>
          <Routes>
            <Route path="/" element={<Root />}>
              <Route index element={<Home />} />
              <Route path="routes/new" element={<NewRoute />} />
              <Route path="about" element={<About />} />
              <Route path="account" element={<Account />} />
              <Route path="login" element={<Login />} />
              <Route path="register" element={<Register />} />
              <Route path="logout" element={<Logout />} />
              <Route path="*" element={<NotFound />} />
            </Route>
          </Routes>
        </BrowserRouter>
      </ToastBarProvider>
    </QueryClientProvider>
  </StrictMode>
);
