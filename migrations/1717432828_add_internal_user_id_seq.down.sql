DROP INDEX IF EXISTS users_useridnum_idx;
ALTER TABLE users
    DROP COLUMN user_id_num;
DROP SEQUENCE IF EXISTS users_user_id_num_seq;
