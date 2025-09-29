CREATE TABLE students (
                          id SERIAL PRIMARY KEY,
                          name VARCHAR(100) NOT NULL,
                          date_of_birth DATE NOT NULL,
                          study_year INT NOT NULL CHECK (study_year >= 1 AND study_year <= 6)
);
