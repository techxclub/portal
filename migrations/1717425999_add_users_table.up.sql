CREATE TABLE IF NOT EXISTS users
(
    user_id             uuid PRIMARY KEY                  DEFAULT uuid_generate_v4(),
    created_time        TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    status              VARCHAR(100)             NOT NULL,
    name                VARCHAR(100),
    phone_number        VARCHAR(20) UNIQUE,
    personal_email      VARCHAR(255)             NOT NULL UNIQUE,
    company_id          INTEGER                  NOT NULL,
    company_name        VARCHAR(100)             NOT NULL,
    role                VARCHAR(255),
    years_of_experience FLOAT,
    work_email          VARCHAR(255)             NOT NULL UNIQUE,
    linkedIn            VARCHAR(255) UNIQUE
);

CREATE INDEX IF NOT EXISTS users_userid_idx ON users (user_id);
CREATE INDEX IF NOT EXISTS users_phone_idx ON users (phone_number);
CREATE INDEX IF NOT EXISTS users_personalemail_idx ON users (personal_email);
CREATE INDEX IF NOT EXISTS users_workemail_idx ON users (work_email);
CREATE INDEX IF NOT EXISTS users_createdtime_idx ON users (created_time);
