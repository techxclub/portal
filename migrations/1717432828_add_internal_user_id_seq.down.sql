DROP INDEX IF EXISTS users_useridnum_idx;
DROP SEQUENCE IF EXISTS users_user_id_num_seq;
ALTER TABLE users DROP COLUMN user_id_num;
