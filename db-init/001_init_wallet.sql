CREATE TABLE IF NOT EXISTS wallets (
                                       id UUID PRIMARY KEY,
                                       balance BIGINT NOT NULL
);

INSERT INTO wallets (id, balance) VALUES
    ('cb7a4b44-8dc2-4fec-9277-61b1fcb5a264', 1000);
