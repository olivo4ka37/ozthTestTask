CREATE TABLE posts
(
    id               SERIAL PRIMARY KEY,
    title            VARCHAR(255) NOT NULL,
    author           VARCHAR(100) NOT NULL,
    content          TEXT         NOT NULL,
    comments_enabled BOOLEAN DEFAULT TRUE
);

CREATE TABLE comments
(
    id        SERIAL PRIMARY KEY,
    post_id   INT REFERENCES posts (id) ON DELETE CASCADE,
    author    VARCHAR(100)  NOT NULL,
    content   VARCHAR(2000) NOT NULL,
    parent_id INT REFERENCES comments (id) ON DELETE CASCADE
);