DROP INDEX IF EXISTS referrals_id_idx;
DROP INDEX IF EXISTS referrals_uuid_idx;
DROP INDEX IF EXISTS referrals_requester_userid_idx;
DROP INDEX IF EXISTS referrals_provider_userid_idx;
DROP INDEX IF EXISTS referrals_create_time_idx;
DROP INDEX IF EXISTS referrals_update_time_idx;
DROP TABLE IF EXISTS referrals;
DROP SEQUENCE IF EXISTS referrals_id_num_seq;
