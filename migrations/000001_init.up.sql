CREATE SCHEMA IF NOT EXISTS nix;

CREATE TABLE IF NOT EXISTS nix.users
(
    id            serial       not null unique PRIMARY KEY,
    username      varchar(255) not null unique,
    email         varchar(255) not null unique
        CONSTRAINT proper_email CHECK (email ~* '^[A-Za-z0-9._%-]+@[A-Za-z0-9.-]+[.][A-Za-z]+$'),
    password_hash varchar(255)
);

CREATE TABLE IF NOT EXISTS nix.posts
(
    id     serial       not null unique PRIMARY KEY,
    UserId integer      not null,
    Title  varchar(255) not null,
    Body   text         not null,
    FOREIGN KEY (UserId) REFERENCES nix.users (id) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS nix.comments
(
    id     serial  not null unique PRIMARY KEY,
    PostId integer not null,
    UserId integer not null,
    Body   text    not null,
    FOREIGN KEY (PostId) REFERENCES nix.posts (id) ON DELETE CASCADE,
    FOREIGN KEY (UserId) REFERENCES nix.users (id) ON DELETE NO ACTION
);

