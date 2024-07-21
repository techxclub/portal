CREATE SEQUENCE IF NOT EXISTS users_user_number_seq START with 1;

CREATE TABLE IF NOT EXISTS users
(
    user_number           BIGINT                   NOT NULL DEFAULT NEXTVAL('users_user_number_seq'),
    user_uuid             uuid PRIMARY KEY                  DEFAULT uuid_generate_v4(),
    create_time           TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time           TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    status                VARCHAR(100)             NOT NULL,
    registered_email      VARCHAR(255)             NOT NULL UNIQUE,
    name                  VARCHAR(100),
    phone_number          VARCHAR(20),
    profile_picture       VARCHAR(1000),
    linkedin              VARCHAR(255),
    gender                VARCHAR(10),
    company_id            INTEGER,
    company_name          VARCHAR(100),
    work_email            VARCHAR(255),
    designation           VARCHAR(100),
    years_of_experience   FLOAT,
    google_auth_details   JSONB,
    technical_information JSONB,
    mentor_config         JSONB
);

CREATE INDEX IF NOT EXISTS users_user_number_idx ON users (user_number);
CREATE INDEX IF NOT EXISTS users_user_uuid_idx ON users (user_uuid);
CREATE INDEX IF NOT EXISTS users_create_time_idx ON users (create_time);
CREATE INDEX IF NOT EXISTS users_update_time_idx ON users (update_time);
CREATE INDEX IF NOT EXISTS users_registered_email_idx ON users (registered_email);
