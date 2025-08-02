CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS games (
    id UUID PRIMARY KEY,
    fen VARCHAR(255) NOT NULL,
    black_player_id UUID REFERENCES users(id),
    white_player_id UUID REFERENCES users(id)
);
