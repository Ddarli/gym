CREATE TABLE users
(
    id           SERIAL PRIMARY KEY,
    username     VARCHAR(255) NOT NULL,
    password     VARCHAR(255) NOT NULL,
    email        VARCHAR(255) NOT NULL,
    phone_number VARCHAR(20)
);

DROP TABLE IF EXISTS users;

INSERT INTO users (username, email, phone_number, password)
VALUES ('John', 'e@mail.com', '+79776453424', 'password123'),
       ('Jane', 'jane@mail.com', '+79776453425', 'password123'),
       ('Alice', 'alice@mail.com', '+79776453426', 'password123'),
       ('Bob', 'bob@mail.com', '+79776453427', 'password123'),
       ('Charlie', 'charlie@mail.com', '+79776453428', 'password123'),
       ('David', 'david@mail.com', '+79776453429', 'password123'),
       ('Eve', 'eve@mail.com', '+79776453430', 'password123'),
       ('Frank', 'frank@mail.com', '+79776453431', 'password123'),
       ('Grace', 'grace@mail.com', '+79776453432', 'password123'),
       ('Hannah', 'hannah@mail.com', '+79776453433', 'password123');


DELETE
FROM users
WHERE username IN (
                   'John', 'Jane', 'Alice', 'Bob', 'Charlie',
                   'David', 'Eve', 'Frank', 'Grace', 'Hannah'
    );