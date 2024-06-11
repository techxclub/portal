CREATE TABLE users
(
    user_id             uuid PRIMARY KEY                  DEFAULT uuid_generate_v4(),
    name                VARCHAR(100),
    company             VARCHAR(100)             NOT NULL,
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
CREATE INDEX IF NOT EXISTS users_createdtime_idx ON users(created_time);
