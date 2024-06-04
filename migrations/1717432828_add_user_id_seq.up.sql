CREATE SEQUENCE users_user_id_seq START with 1;
ALTER TABLE users
    ADD COLUMN user_id BIGINT NOT NULL DEFAULT NEXTVAL('users_user_id_seq');
CREATE INDEX IF NOT EXISTS users_userid_idx ON users(user_id);
