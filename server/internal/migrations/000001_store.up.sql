create table if not exists users_auth (
    record_id serial,
    record_type text,
    user_id varchar(27) references keeper_auth (user_id) not null,
    description text,
    metadata text,
    user_encrypted_login varchar(255),
    user_encrypted_password varchar(255),
    del_flag boolean default false,
    created_at timestamp not null default now()
);

create table if not exists users_files (
    record_id serial,
    user_id varchar(27) references keeper_auth (user_id) not null,
    description text,
    metadata text,
    user_file bytea,
    del_flag boolean default false,
    created_at timestamp not null default now()
);

create table if not exists users_cards (
    record_id serial,
    user_id varchar(27) references keeper_auth (user_id) not null,
    description text,
    metadata text,
    encrypted_card text,
    del_flag boolean default false,
    created_at timestamp not null default now()
);
