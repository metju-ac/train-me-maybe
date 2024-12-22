# Regiojet API

The API is not public, but since their website uses it, it is possible to reverse-engineer it.

- See [official documentation](https://regiojet.com/about-us/affiliate/api).
- See [swagger documentation](https://app.swaggerhub.com/apis/regiojet/affiliate/1.1.0).
- See [downloaded swagger](./regiojet-affiliate-1.1.0-resolved.json).

We use [Bruno](https://www.usebruno.com/) for storing the API definiton. The folder `./bruno`
is synchronized with the Bruno project and is versioned in git.

All of the example requests are already configured in the Bruno project.

## Using Bruno

1. Inside Bruno, open a collection (the folder `./bruno`).
2. Choose an environment (upper right-hand corner) as `Regiojet public API`.
3. Fill in secret values like this:

![Bruno - where to configure variables](./assets/bruno-configure-variables.png)

![Bruno - configuring variables](./assets/bruno-configure-variables-2.png)


## Basic description of endpoints

- `GET /routes/search/simple` - used when searching by origin and destination + date:
```
https://brn-ybus-pubapi.sa.cz/restapi/routes/search/simple?tariffs=CZECH_STUDENT_PASS_26&toLocationType=CITY&toLocationId=10202003&fromLocationType=CITY&fromLocationId=10202002&departureDate=2024-12-22&fromLocationName=&toLocationName=
```

- TODO: study the normal flow of searching and describe it here

## Authentication

The downloaded swagger is missing the authentication process. Therefore, we added it manually
as the `Login with kreditová jízdenka` request.

- `POST /users/login/registeredAccount`
```json
{"accountCode":"NUMBER","password":"PASSWORD"}
```

After logging in successfully, we are granted a UUID. 
This should be used as a `Authorization: Bearer <uuid_token>` header in subsequent requests.
