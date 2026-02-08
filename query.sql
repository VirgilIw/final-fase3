CREATE TABLE account (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
);

-- CREATE TYPE user_role AS ENUM ('admin', 'user');


-- ALTER TABLE account
-- ADD COLUMN role user_role;

-- ALTER TABLE pengguna ALTER COLUMN nama TYPE VARCHAR(100);

SELECT * from account;
CREATE TABLE users (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    account_id INT UNIQUE NOT NULL REFERENCES account(id),
    user_name VARCHAR(255),
    user_image VARCHAR(255),
    user_bio VARCHAR(255)
);

-- ALTER TABLE users
-- ADD COLUMN created_at TIMESTAMP NOT NULL DEFAULT NOW(),
-- ADD COLUMN updated_at TIMESTAMP NOT NULL DEFAULT NOW();

CREATE TABLE posts (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id),
    post_text VARCHAR(255),
    post_image VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE like_post (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id),
    post_id INT NOT NULL REFERENCES posts(id),
    UNIQUE(user_id, post_id)
);

CREATE TABLE comment_post (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id),
    post_id INT NOT NULL REFERENCES posts(id),
    new_comment VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE follow_user (
    follower_id INT REFERENCES users(id),
    following_id INT REFERENCES users(id),
    UNIQUE(follower_id, following_id)
);

CREATE TABLE notification (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id INT REFERENCES users(id),
    from_user_id INT REFERENCES users(id),
    post_id INT REFERENCES posts(id),
    notification_type VARCHAR(50),
    created_at TIMESTAMP DEFAULT NOW()
);
