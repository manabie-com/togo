CREATE TABLE IF NOT EXISTS member (
    -- id
    id INT GENERATED ALWAYS AS IDENTITY,
    -- email
    email VARCHAR(255) NOT NULL,
    -- name
    name VARCHAR(256) NOT NULL,
    -- create_at
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    -- update_at
    updated_at TIMESTAMP WITHOUT TIME ZONE,
    -- pk
    PRIMARY KEY (id)
);
