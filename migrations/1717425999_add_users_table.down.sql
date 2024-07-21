DROP INDEX IF EXISTS users_user_number_idx;
DROP INDEX IF EXISTS users_user_uuid_idx;
DROP INDEX IF EXISTS users_create_time_idx;
DROP INDEX IF EXISTS users_update_time_idx;
DROP INDEX IF EXISTS users_registered_email_idx;
DROP TABLE IF EXISTS users;
DROP SEQUENCE IF EXISTS users_user_number_seq;
