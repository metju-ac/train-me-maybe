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


## Description of endpoints

- Look at the Booking API section of [the docs](https://regiojet.com/about-us/affiliate/api).

## Authentication

The downloaded swagger is missing the authentication process. Therefore, we added it manually
as the `Login with kreditová jízdenka` request.

- `POST /users/login/registeredAccount`
```json
{"accountCode":"NUMBER","password":"PASSWORD"}
```

After logging in successfully, we are granted a UUID. 
This should be used as a `Authorization: Bearer <uuid_token>` header in subsequent requests.

## Generated API client

How we have generated the code inside the `generated-api-client-go` folder:

1. `npm install @openapitools/openapi-generator-cli -g`
2. `openapi-generator-cli generate -i ./regiojet-affiliate-1.1.0-resolved.json -g go -o generated-api-client-go --skip-validate-spec`

Look inside [the readme](./generated-api-client-go/README.md) to see how to use the generated client.
