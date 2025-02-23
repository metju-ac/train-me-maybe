CREATE TABLE notifications
(
    id                    SERIAL PRIMARY KEY,
    user_email            VARCHAR(255),
    from_station_id       BIGINT         NOT NULL,
    to_station_id         BIGINT         NOT NULL,
    route_id              VARCHAR(255)   NOT NULL,
    selected_seat_classes TEXT[]         NOT NULL,
    notification_time     TIMESTAMP      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    seat_class            VARCHAR(255),
    FOREIGN KEY (user_email) REFERENCES users (email)
);