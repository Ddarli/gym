CREATE TABLE users
(
    id       SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL
);

CREATE TABLE classes
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(255) NOT NULL
);

CREATE TABLE bookings
(
    id                 SERIAL PRIMARY KEY,
    user_id            INT NOT NULL,
    scheduled_class_id INT NOT NULL,
    booking_time       TIMESTAMP,
    status             INT NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_class FOREIGN KEY (scheduled_class_id) REFERENCES classes(id) ON DELETE CASCADE
);

INSERT INTO users(username)
VALUES ('John'), ('Jane'), ('Alice'), ('Bob'), ('Charlie'),
       ('David'), ('Eve'), ('Frank'), ('Grace'), ('Hannah');

INSERT INTO classes(name)
VALUES ('Yoga'),
       ('Pilates'),
       ('Zumba'),
       ('Spinning'),
       ('Boxing'),
       ('Karate'),
       ('Crossfit'),
       ('Stretching'),
       ('Swimming'),
       ('Running');

INSERT INTO bookings(user_id, scheduled_class_id, booking_time, status)
VALUES (1, 1, NOW(), 1),
       (2, 2, NOW(), 2),
       (3, 3, NOW(), 1),
       (4, 4, NOW(), 1),
       (5, 5, NOW(), 2),
       (6, 6, NOW(), 3),
       (7, 7, NOW(), 1),
       (8, 8, NOW(), 2),
       (9, 9, NOW(), 3),
       (10, 10, NOW(), 1);




TRUNCATE TABLE users;
TRUNCATE TABLE classes;
TRUNCATE TABLE bookings;
DROP TABLE users;
DROP TABLE classes;
DROP TABLE bookings;