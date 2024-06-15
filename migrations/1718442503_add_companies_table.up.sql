CREATE SEQUENCE companies_id_num_seq START with 1;

CREATE TABLE companies
(
    id               INTEGER PRIMARY KEY NOT NULL DEFAULT NEXTVAL('users_user_id_num_seq'),
    name             VARCHAR(100)        NOT NULL,
    small_logo       VARCHAR(1000),
    big_logo         VARCHAR(1000),
    official_website VARCHAR(1000),
    careers_page     VARCHAR(1000),
    priority         INTEGER,
    verified         BOOLEAN             NOT NULL DEFAULT FALSE,
    popular          BOOLEAN             NOT NULL DEFAULT FALSE
);

CREATE INDEX IF NOT EXISTS companies_id_idx ON companies(id);
CREATE INDEX IF NOT EXISTS companies_name_idx ON companies(name);
