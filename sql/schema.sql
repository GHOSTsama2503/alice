CREATE TABLE users (
	id INTEGER NOT NULL,
	language_code TEXT NOT NULL
);


CREATE TABLE groups (
	id INTEGER PRIMARY KEY,
	title TEXT NOT NULL,
	is_private INTEGER NOT NULL
);


CREATE TABLE messages (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	group_id INTEGER NOT NULL,
	welcome_message TEXT NOT NULL DEFAULT '',

	CONSTRAINT fk_messages_group_id FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE
);


CREATE TABLE groups_settings (
	id INTEGER PRIMARY KEY,
	group_id INTEGER NOT NULL,
	is_welcome_message_enabled INTEGER NOT NULL,

	CONSTRAINT fk_group_settings_group_id FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE
);
