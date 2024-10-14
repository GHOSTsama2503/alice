drop table if exists chat_triggers;
drop table if exists chat_settings;
drop trigger if exists update_chats_updated_at on chats;
drop table if exists chats;
drop type if exists ChatType;
drop function if exists update_updated_at();
