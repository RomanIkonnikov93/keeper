create table if not exists keeper_auth (
    user_id varchar(27) unique not null,
    user_login varchar(255) unique not null,
    user_password varchar(255) not null,
    created_at timestamp not null default now()
);
