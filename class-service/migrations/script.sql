CREATE TABLE classes (
                         id SERIAL PRIMARY KEY,
                         name VARCHAR(100) NOT NULL,
                         description TEXT,
                         capacity INT NOT NULL
);


INSERT INTO classes (name, description, capacity)
VALUES
    ('Yoga', 'A relaxing yoga class for all levels.', 20),
    ('Pilates', 'Core strengthening and flexibility.', 15),
    ('HIIT', 'High-Intensity Interval Training workout.', 25),
    ('Zumba', 'Dance fitness class.', 30),
    ('Spin', 'Indoor cycling class.', 20),
    ('Boxing', 'Introduction to boxing techniques.', 10),
    ('Crossfit', 'Strength and conditioning workout.', 15),
    ('Stretching', 'Improve flexibility and relaxation.', 20),
    ('Meditation', 'Guided meditation for mindfulness.', 25),
    ('Strength Training', 'Weight lifting and strength training.', 15);