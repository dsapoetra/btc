CREATE TABLE IF NOT EXIST transactions (
    amount float NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW ()
);

-- Add indexes
CREATE INDEX created_at_idx ON transactions (created_at);