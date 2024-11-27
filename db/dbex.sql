CREATE
EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE songs
(
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title           VARCHAR(50) NOT NULL,
    group_song      VARCHAR(50) NOT NULL,
    release_date    DATE NOT NULL,
    text_song       TEXT NOT NULL,
    link_song       TEXT NOT NULL
);

CREATE INDEX idx_id ON songs(id);
