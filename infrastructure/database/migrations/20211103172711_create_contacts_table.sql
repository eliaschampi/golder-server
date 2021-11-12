-- migrate:up
CREATE TABLE IF NOT EXISTS contacts
(
    code        uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name        VARCHAR(255) NOT NULL,
    type        BPCHAR(1)    NOT NULL,
    email       VARCHAR(50),
    phone       VARCHAR(30),
    address     VARCHAR(120),
    description VARCHAR(300),
    created_at  timestamp(6)     DEFAULT CURRENT_TIMESTAMP(6),
    updated_at  timestamp(6)     DEFAULT CURRENT_TIMESTAMP(6)
);

-- migrate:down
DROP TABLE IF EXISTS contacts;
