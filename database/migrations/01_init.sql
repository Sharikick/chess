CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS games (
    id UUID PRIMARY KEY,
    fen VARCHAR(255) NOT NULL,
    black_player_id INT REFERENCES users(id),
    white_player_id INT REFERENCES users(id),
    current_turn VARCHAR(5) CHECK(current_turn IN('white', 'black')),
    CONSTRAINT defferent_players CHECK (black_player_id != white_player_id)
);
