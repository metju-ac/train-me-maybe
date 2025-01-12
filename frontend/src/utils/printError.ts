import ToastBarContext from "@components/ToastBarContext";
import { AxiosError } from "axios";

type Decondextify<T> = T extends { Provider: React.Provider<infer P> }
  ? P
  : never;

export default function printError(
  setToast: Decondextify<typeof ToastBarContext>["setToast"],
  err: Error,
  prefix = "Error"
) {
  /* eslint-disable */
  const errMessage =
    (err as AxiosError<any>).response?.data?.error ??
    (err as AxiosError<any>).response?.data ??
    err.message;
  /* eslint-enable */

  setToast(prefix + ": " + errMessage, "error");
}
