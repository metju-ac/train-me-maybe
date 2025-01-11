# Project architecture

The regiojet watchdog can be run in two modes:

- Single-user mode
- Multi-user mode

## Single-user mode

We assume that the user is experienced with the command line and can configure and run the application,
assisted with readme instructions.

- All configuration is done via `config.toml` (lower priority) or environemnt variables (higher priority).
- User typically runs this on their own machine/server
- No support for multiple users - the user can only enter one Regiojet account (Kreditová jízdenka)
- No support for multiple routes - the user can only watch one route
  - if desired, the user can run multiple instances of the application
- The application exists after it finds a route with a free seat and sends a notification (or purchases the ticket automatically)

## Multi-user mode

This mode is harder to set up and maintain, but it allows multiple inexperienced users to "just use" the application.
To avoid the need for a database, the application is designed to store its state fully in-memory.
However, if the application crashes, all watched routes are of course lost.

- There is a web interface for end-users.
- Users have to register themselves.
- Users log in using the email (which should receive the notifications) and password.
- Once logged in, users can configure a single route to watch.
- If the user wants, they can also opt-in for "automatic purchase" of the ticket via a checkbox.
  - In this case, they have to fill in the Regiojet account (Kreditová jízdenka) credentials.
  - These credentials are first validated by the application, and then stored in the database.
- When a user submits the route configuration, a goroutine is spawned to periodically watch the route.
  - This goroutine is exactly the same function as the one in the single-user mode.
- The user can enter another route to watch.
- The user sees all watched routes (for his account only) and can cancel watching them.
  - To support cancellation of the route watching, the user can click on a button "Stop watching".
    - Technically, this is achieved by cancelling the goroutine (calling `cancel()` on a context; the goroutine must watch for this cancellation)
    - In-memory, for each user we must store the list of watched routes and their respective contexts.
- The user has these config options:
  - purchase cutoff time - how much time before the departure should the ticket be purchased
  - jmeno a heslo do kreditove jizdenky
  - castka, kdyz mam mensi kredit nez tohle, tak mi posli email
  - default tariff class

### DB design

Watch out - every int should be (int64) in Go.

- Table `User` for registered user

  - email (PK)
  - password hash (password hash)
  - vielleicht salz
  - OPTIONAL - cislo kreditove jizdenky
  - OPTIONAL - heslo kreditove jizdenky (plaintext)
  - OPTIONAL - cut-off time
  - OPTIONAL - minimalni kredit
  - OPTIONAL - default tariff key (string)
  - mozna created_at
  - mozna updated_at
  - mozna deleted_at
  - kolik nam dluzi piv

- Table `WatchedRoute` for currently watched routes - so that re-creation of the watched routes is possible after a crash.
  When a free seat is found, a row is deleted from this table. The user can also delete the row manually (must also cancel the goroutine).

  - ID (PK)
  - user_email (FK)
  - from station ID
  - to station ID
  - route ID
  - selected seat classes (LOW_COST atd)
  - auto purchase (bool)
  - OPTIONAL - tariff class (e.g. STUDENT, SENIOR, ADULT)
  - OPTIONAL - credit user
  - OPTIONAL - credit password
  - OPTIONAL - cut-off time
  - OPTIONAL - minimal credit

- OPTIONAL - Table `SuccessfulPurchase` for successful purchases
  - ID (PK)
  - user_email (FK)
  - from station ID
  - to station ID
  - route ID
  - tariff class
  - selected seat classes
  - purchase time
  - price
  - currency
  - seat class

### API and Front end

All POST bodies are JSON.

- POST /api/register

  - body: email, password
  - returns 200 OK, if easy to implement, also automatically logs in the user

- POST /api/login
  - body: email, password
  - returns 200 OK with Set-Cookie header, inside which is probably a JWT token, signed/MACked by the server (secret key should be envvar or config)

=== BEGIN endpoints accessible only for authorized users (probably middleware) ===

- GET /api/auth/station

  - returns a list of all stations (locations mapped to our DTO)
  - sets the correct HTTP headers so that the browser caches this response for a long time
  - the server caches these locations in-memory (probably on startup)

- GET /api/auth/route

  - URL params: fromStationId, toStationId, date (date only, without time)
  - returns routes for the given date
  - also sets HTTP headers for browser caching, probably not necessary to cache on server

- GET /api/seatClass

  - NOT needed - we will hardcode this on the front end

- GET /api/tariffClass

  - NOT needed - we will hardcode this on the front end

- POST /api/auth/watchedRoute

  - body: autoPurchase (bool), fromStationId (int64), toStationId (int64), routeId (string), tariffClass (string), selectedSeatClasses, creditUser (optional string), creditPassword (optional string), cutOffTime (optional int), minimalCredit (optional int)
    - if autoPurchase is true, creditUser and creditPassword and tarriff class must be filled in

- GET /api/auth/user

  - gets details about the current logged in user (from the cookie)

- PUT /api/auth/user

  - body: cutOffTime (int), minimalCredit (int), creditUser (string), creditPassword (string), tarriffClass (string)
  - modifies the current logged in user (from the cookie)

- GET /api/auth/watchedRoute

  - returns all watched routes for the current user

- DELETE /api/auth/watchedRoute/:id
  - deletes the watched route with the given ID
