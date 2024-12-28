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
- Users do not have to register themselves.
- Users log in using just the email (no password) which should receive the notifications. 
- Once logged in, users can configure a single route to watch.
- If the user wants, they can also opt-in for "automatic purchase" of the ticket via a checkbox.
  - In this case, they have to fill in the Regiojet account (Kreditová jízdenka) credentials.
- When a user submits the route configuration, a goroutine is spawned to periodically watch the route.
  - This goroutine is exactly the same function as the one in the single-user mode.
- The user can enter another route to watch.
- To support cancellation of the route watching, the user can click on a button "Stop watching".
  - Technically, this is achieved by cancelling the goroutine (calling `cancel()` on a context; the goroutine must watch for this cancellation)
  - In-memory, for each user we must store the list of watched routes and their respective contexts.  
