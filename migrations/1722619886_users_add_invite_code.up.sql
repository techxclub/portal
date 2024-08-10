ALTER TABLE users
    ADD COLUMN invite_code VARCHAR(255) UNIQUE;

CREATE INDEX users_invite_code_idx ON users (invite_code);
