create table if not exists users (
    id serial primary key,
    email text not null unique,
    pass_hash text not null
);

create index if not exists idx_email on users(email);

create table if not exists apps (
    id integer primary key,
    name text not null unique,
    secret text not null unique
);

create table if not exists admins (
    id integer primary key
);