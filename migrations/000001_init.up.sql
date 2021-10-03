CREATE SCHEMA IF NOT EXISTS nix;

CREATE TABLE IF NOT EXISTS nix.posts
(
    id     BIGINT       not null unique,
    UserId BIGINT       not null,
    Title  varchar(255) not null,
    Body   text         not null
);


CREATE TABLE IF NOT EXISTS nix.comments
(
    id     BIGINT       not null unique,
    PostId BIGINT       not null,
    name   varchar(255) not null,
    email  varchar(255) not null,
    Body   text         not null,
    FOREIGN KEY (PostId) REFERENCES nix.posts (id)
);
