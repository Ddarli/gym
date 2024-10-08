CREATE TABLE trainers
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE classes
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

INSERT INTO classes(name)
VALUES ('Yoga'),
       ('Pilates'),
       ('HIIT'),
       ('Zumba'),
       ('Spin'),
       ('Boxing'),
       ('Crossfit'),
       ('Stretching'),
       ('Meditation'),
       ('Strength Training');

INSERT INTO trainers(name)
VALUES ('Ivanov Ivan'),
       ('Petrov Petr'),
       ('Sidorova Svetlana'),
       ('Kuznetsov Alexey'),
       ('Smirnova Olga'),
       ('Fedorova Anna'),
       ('Yegorov Roman'),
       ('Soloviev Nikolai'),
       ('Morozova Ksenia'),
       ('Tikhonov Dmitry');


CREATE TABLE schedules
(
    id         SERIAL PRIMARY KEY,
    class_id   INT       NOT NULL REFERENCES classes (id),
    trainer_id INT       NOT NULL REFERENCES trainers (id),
    start_time TIMESTAMP NOT NULL,
    end_time   TIMESTAMP NOT NULL
);


INSERT INTO schedules (class_id, trainer_id, start_time, end_time)
VALUES (1, 1, '2024-10-07 08:00:00', '2024-10-07 09:00:00'),
       (2, 2, '2024-10-07 09:00:00', '2024-10-07 10:00:00'),
       (3, 3, '2024-10-07 10:00:00', '2024-10-07 11:00:00'),
       (4, 4, '2024-10-07 11:00:00', '2024-10-07 12:00:00'),
       (5, 5, '2024-10-07 12:00:00', '2024-10-07 13:00:00'),
       (6, 6, '2024-10-07 13:00:00', '2024-10-07 14:00:00'),
       (7, 7, '2024-10-07 14:00:00', '2024-10-07 15:00:00'),
       (8, 8, '2024-10-07 15:00:00', '2024-10-07 16:00:00'),
       (9, 9, '2024-10-07 16:00:00', '2024-10-07 17:00:00'),
       (10, 10, '2024-10-07 17:00:00', '2024-10-07 18:00:00');


TRUNCATE TABLE schedules;
DROP TABLE schedules;
