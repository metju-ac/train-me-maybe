import { createContext } from "react";

const ToastBarContext = createContext({
  // eslint-disable-next-line @typescript-eslint/no-unused-vars, @typescript-eslint/no-empty-function
  setToast: (_message: string, _severity: "success" | "error") => {},
});

export default ToastBarContext;
