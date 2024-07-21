CREATE SEQUENCE IF NOT EXISTS referrals_id_num_seq START with 1;

CREATE TABLE IF NOT EXISTS referrals
(
    id                INTEGER PRIMARY KEY      NOT NULL DEFAULT NEXTVAL('referrals_id_num_seq'),
    uuid              uuid                     NOT NULL DEFAULT uuid_generate_v4(),
    create_time       TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time       TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    requester_user_id VARCHAR(255)             NOT NULL,
    provider_user_id  VARCHAR(255)             NOT NULL,
    company_id        INTEGER                  NOT NULL,
    job_link          VARCHAR(255),
    status            VARCHAR(100)             NOT NULL
);

CREATE INDEX IF NOT EXISTS referrals_id_idx ON referrals (id);
CREATE INDEX IF NOT EXISTS referrals_create_time_idx ON referrals (create_time);
CREATE INDEX IF NOT EXISTS referrals_update_time_idx ON referrals (update_time);
CREATE INDEX IF NOT EXISTS referrals_uuid_idx ON referrals (uuid);
CREATE INDEX IF NOT EXISTS referrals_requester_userid_idx ON referrals (requester_user_id);
CREATE INDEX IF NOT EXISTS referrals_provider_userid_idx ON referrals (provider_user_id);
