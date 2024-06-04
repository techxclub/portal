DROP INDEX IF EXISTS users_userid_idx;
DROP SEQUENCE IF EXISTS users_user_id_seq;
ALTER TABLE users DROP COLUMN user_id;
