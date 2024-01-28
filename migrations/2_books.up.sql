CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    published_date VARCHAR(255),
    original_language VARCHAR(255)
);

INSERT INTO books (title, author, published_date, original_language) 
VALUES ('7 Habits of Highly Effective People', 'Stephen Covey', '1989-08-15', 'English');

