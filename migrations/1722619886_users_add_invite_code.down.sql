DROP INDEX IF EXISTS users_invite_code_idx;
ALTER TABLE users
    DROP COLUMN IF EXISTS invite_code;
