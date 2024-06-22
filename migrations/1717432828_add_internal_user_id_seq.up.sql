CREATE SEQUENCE IF NOT EXISTS users_user_id_num_seq START with 1;
ALTER TABLE users
    ADD COLUMN user_id_num BIGINT NOT NULL DEFAULT NEXTVAL('users_user_id_num_seq');
CREATE INDEX IF NOT EXISTS users_useridnum_idx ON users (user_id_num);
