CREATE TABLE IF NOT EXISTS users (
	id UUID NOT NULL DEFAULT uuid_generate_v1(),
	name TEXT NOT NULL DEFAULT '',
	age INTEGER NOT NULL DEFAULT 0,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT current_timestamp,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT current_timestamp,
	deleted_at TIMESTAMP WITH TIME ZONE NULL,
	profile_id UUID NOT NULL,
	PRIMARY KEY (id),
	FOREIGN KEY (profile_id) REFERENCES profiles(id)
);
