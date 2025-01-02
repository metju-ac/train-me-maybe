import AddWatchedRoute from "@components/AddWatchedRoute";
import WatchedRoutesTable from "@components/WatchedRoutesTable";
import useIsLoggedIn from "@utils/useIsLoggedIn";

export default function Home() {
  const isLoggedIn = useIsLoggedIn();

  if (!isLoggedIn) {
    return (
      <>
        <h1>You are not logged in</h1>
        <p>To use the app you need to register.</p>
      </>
    );
  }

  return (
    <>
      <AddWatchedRoute />
      <WatchedRoutesTable />
    </>
  );
}
