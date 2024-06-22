DROP INDEX IF EXISTS referrals_id_idx;
DROP INDEX IF EXISTS referrals_requester_userid_idx;
DROP INDEX IF EXISTS referrals_provider_userid_idx;
DROP INDEX IF EXISTS referrals_createdtime_idx;
DROP TABLE IF EXISTS referrals;
DROP SEQUENCE IF EXISTS referrals_id_num_seq;
