CREATE TABLE users
(
    user_id             VARCHAR(50)              NOT NULL UNIQUE PRIMARY KEY,
    first_name          VARCHAR(100),
    last_name           VARCHAR(100),
    years_of_experience FLOAT,
    personal_email      VARCHAR(255)             NOT NULL UNIQUE,
    work_email          VARCHAR(255)             NOT NULL UNIQUE,
    phone_number        VARCHAR(20) UNIQUE,
    linkedIn            VARCHAR(255) UNIQUE,
    role                VARCHAR(100),
    created_time        TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS users_userid_idx ON users(user_id);
CREATE INDEX IF NOT EXISTS users_phone_idx ON users(phone_number);
CREATE INDEX IF NOT EXISTS users_personalemail_idx ON users(personal_email);
CREATE INDEX IF NOT EXISTS users_workemail_idx ON users(work_email);
CREATE INDEX IF NOT EXISTS users_createdat_idx ON users(created_time);
