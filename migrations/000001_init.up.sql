CREATE TABLE users
(
    id                   SERIAL PRIMARY KEY,
    email                VARCHAR(255) UNIQUE NOT NULL,
    password_hash        VARCHAR(255)        NOT NULL,
    salt                 VARCHAR(255)        NOT NULL,
    credit_user_number   VARCHAR(255),
    credit_user_password VARCHAR(255),
    cut_off_time         INT,
    minimal_credit       INT,
    default_tariff_class VARCHAR(255),
    beers_owed           DECIMAL(10, 2)      NOT NULL DEFAULT 0,
    created_at           TIMESTAMP,
    updated_at           TIMESTAMP,
    deleted_at           TIMESTAMP
);

CREATE TABLE watched_routes
(
    id                    SERIAL PRIMARY KEY,
    user_email            VARCHAR(255),
    from_station_id       BIGINT        NOT NULL,
    to_station_id         BIGINT        NOT NULL,
    route_id              VARCHAR(255)  NOT NULL,
    selected_seat_classes TEXT[]        NOT NULL,
    auto_purchase         BOOLEAN       NOT NULL DEFAULT FALSE,
    tariff_class          VARCHAR(255),
    credit_user_number    VARCHAR(255),
    credit_user_password  VARCHAR(255),
    cut_off_time          INT,
    minimal_credit        INT,
    FOREIGN KEY (user_email) REFERENCES users (email)
);

CREATE TABLE successful_purchases
(
    id                    SERIAL PRIMARY KEY,
    user_email            VARCHAR(255),
    from_station_id       BIGINT         NOT NULL,
    to_station_id         BIGINT         NOT NULL,
    route_id              VARCHAR(255)   NOT NULL,
    tariff_class          VARCHAR(255)   NOT NULL,
    selected_seat_classes TEXT[]         NOT NULL,
    purchase_time         TIMESTAMP      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    price                 DECIMAL(10, 2) NOT NULL,
    currency              VARCHAR(16),
    seat_class            VARCHAR(255),
    FOREIGN KEY (user_email) REFERENCES users (email)
);