create or replace function update_updated_at() returns trigger as $$ begin
    new.updated_at = current_timestamp;
end; $$ language plpgsql;

drop type if exists ChatType;
create type ChatType as enum ('private', 'group', 'supergroup', 'channel');

create table if not exists chats (
    id          bigint      primary key,
    type        ChatType    not null,
    title       text        not null,
    created_at  timestamp   not null default current_timestamp,
    updated_at  timestamp   not null default current_timestamp
);

create or replace trigger update_chats_updated_at
before update on chats
for each row
execute function update_updated_at();

create table if not exists chat_settings (
    id                          bigint  primary key references chats(id) on update cascade on delete cascade,
    is_captcha_enabled          boolean not null default false,
    is_welcome_message_enabled  boolean not null default false
);

create table if not exists chat_triggers (
    chat_id bigint  primary key references chats(id) on update cascade on delete cascade,
    key     text    not null,
    message text    not null,

    unique (chat_id, key)
);
