create table if not exists users (
    id              bigint      primary key,
    language_code   text,
    created_at      timestamp   not null default current_timestamp,
    updated_at      timestamp   not null default current_timestamp
);

create or replace trigger update_users_updated_at
before update on users
for each row
execute function update_updated_at();

create table if not exists user_stats (
    id          integer primary key generated always as identity,
    chat_id     bigint  not null references chats(id) on update cascade on delete cascade,
    user_id     bigint  not null references users(id) on update cascade on delete cascade,
    warnings    integer not null default 0 check (warnings >= 0),

    unique (chat_id, user_id)
);
