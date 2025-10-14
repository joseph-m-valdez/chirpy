-- +goose Up
CREATE TABLE chirps (
	id				uuid NOT NULL,
	user_id		uuid NOT NULL,
	created_at	timestamp NOT NULL,	
	updated_at	timestamp NOT NULL,	
	body			text			NOT NULL,
	CONSTRAINT chirps_user_id_fkey
		FOREIGN KEY (user_id)
		REFERENCES	users(id)
		ON DELETE CASCADE
);

-- +goose Down
DROP TABLE chirps;
