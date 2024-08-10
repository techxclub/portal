CREATE TABLE invites
(
    id                SERIAL PRIMARY KEY,
    created_time      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    code              VARCHAR(255) NOT NULL,
    invited_user_uuid UUID NOT NULL,
    FOREIGN KEY (invited_user_uuid) REFERENCES users (user_uuid) ON UPDATE CASCADE,
    FOREIGN KEY (code) REFERENCES users (invite_code) ON UPDATE CASCADE
);

CREATE INDEX invites_id_idx ON invites (id);
CREATE INDEX invites_created_time_idx ON invites (created_time);
CREATE INDEX invites_code_idx ON invites (code);
CREATE INDEX invites_invited_user_uuid_idx ON invites (invited_user_uuid);
