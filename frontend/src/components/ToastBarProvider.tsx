import { Alert, Snackbar } from "@mui/material";
import { PropsWithChildren, useState } from "react";
import ToastBarContext from "./ToastBarContext";

export default function ToastBarProvider(
  props: PropsWithChildren<Record<never, never>>
) {
  const [state, setState] = useState({
    open: false,
    message: "",
    severity: "success" as "success" | "error",
  });

  const handleClose = (_?: React.SyntheticEvent | Event, reason?: string) => {
    if (reason === "clickaway") {
      return;
    }
    setState((prev) => ({ ...prev, open: false }));
  };

  return (
    <>
      <Snackbar
        open={state.open}
        autoHideDuration={6000}
        onClose={handleClose}
        anchorOrigin={{ vertical: "top", horizontal: "right" }}
      >
        <Alert
          onClose={handleClose}
          severity={state.severity}
          sx={{ width: "100%" }}
        >
          {state.message}
        </Alert>
      </Snackbar>
      <ToastBarContext.Provider
        value={{
          setToast: (message: string, severity: "success" | "error") => {
            setState({ open: true, message, severity });
          },
        }}
      >
        {props.children}
      </ToastBarContext.Provider>
    </>
  );
}
