import WatchedRoutesTable from "@components/WatchedRoutesTable";
import { getCookie, USER_COOKIE } from "@utils/getCookie";

export default function Home() {
  const isLoggedIn = getCookie(USER_COOKIE) !== null;

  if (!isLoggedIn) {
    return (
      <>
        <h1>You are not logged in</h1>
        <p>To use the app you need to register.</p>
      </>
    );
  }

  return <WatchedRoutesTable />;
}
