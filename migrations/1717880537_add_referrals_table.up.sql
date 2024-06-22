CREATE SEQUENCE IF NOT EXISTS referrals_id_num_seq START with 1;

CREATE TABLE IF NOT EXISTS referrals
(
    id                INTEGER PRIMARY KEY      NOT NULL DEFAULT NEXTVAL('referrals_id_num_seq'),
    requester_user_id VARCHAR(255)             NOT NULL,
    provider_user_id  VARCHAR(255)             NOT NULL,
    company_id        INTEGER                  NOT NULL,
    job_link          VARCHAR(255),
    status            VARCHAR(100)             NOT NULL,
    created_time      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS referrals_id_idx ON referrals (id);
CREATE INDEX IF NOT EXISTS referrals_requester_userid_idx ON referrals (requester_user_id);
CREATE INDEX IF NOT EXISTS referrals_provider_userid_idx ON referrals (provider_user_id);
CREATE INDEX IF NOT EXISTS referrals_createdtime_idx ON referrals (created_time);
