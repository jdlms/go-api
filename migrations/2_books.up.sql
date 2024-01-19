CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY
    title VARCHAR NOT NULL
    author VARCHAR NOT NULL
    published_date DATE
    original_language VARCHAR
)

INSERT INTO books(id) VALUES("1")
INSERT INTO books(title) VALUES("7 Habits of Highly Effective People")
INSERT INTO books(author) VALUES("Stephen Covey")
INSERT INTO books(published_date) VALUES("1989-08-15")
INSERT INTO books(original_language) VALUES("English")
