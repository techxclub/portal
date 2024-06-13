CREATE TABLE referrals
(
    id                uuid PRIMARY KEY                  DEFAULT uuid_generate_v4(),
    requester_user_id VARCHAR(255)             NOT NULL,
    provider_user_id  VARCHAR(255)             NOT NULL,
    company           VARCHAR(255)             NOT NULL,
    job_link          VARCHAR(255),
    status            VARCHAR(100)             NOT NULL,
    created_time      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS referrals_id_idx ON referrals(id);
CREATE INDEX IF NOT EXISTS referrals_requester_userid_idx ON referrals(requester_user_id);
CREATE INDEX IF NOT EXISTS referrals_provider_userid_idx ON referrals(provider_user_id);
CREATE INDEX IF NOT EXISTS referrals_createdtime_idx ON referrals(created_time);
