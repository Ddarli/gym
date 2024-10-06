
CREATE TABLE trainers
(
    id             SERIAL PRIMARY KEY,
    name           VARCHAR(100) NOT NULL,
    specialization VARCHAR(100) NOT NULL
);


INSERT INTO trainers (name, specialization)
VALUES ('Ivanov Ivan', 'Fitness Trainer'),
       ('Petrov Petr', 'Strength Trainer'),
       ('Sidorova Svetlana', 'Yoga Trainer'),
       ('Kuznetsov Alexey', 'Swimming'),
       ('Smirnova Olga', 'CrossFit'),
       ('Fedorova Anna', 'Pilates'),
       ('Yegorov Roman', 'Boxing'),
       ('Soloviev Nikolai', 'Dancing'),
       ('Morozova Ksenia', 'Bodybuilding'),
       ('Tikhonov Dmitry', 'Sports Gymnastics');



DROP TABLE trainers;