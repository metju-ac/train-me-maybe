import './App.css'
import { getCookie, USER_COOKIE } from './utils/getCookie'

function App() {
  const isLoggedIn = getCookie(USER_COOKIE) !== null

  if (!isLoggedIn) {
    return <h1>Not logged in</h1>
  }

  return <h1>Logged in</h1>
}

export default App
