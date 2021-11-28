CREATE TABLE auth."user" (
	id		SERIAL PRIMARY KEY,
	first_name	TEXT,
	last_name	TEXT,
	email		TEXT,
	password 	TEXT,
	created 	TIMESTAMP NOT NULL DEFAULT now()
);

GRANT ALL ON TABLE auth."user" TO finance_calc;
GRANT ALL ON SEQUENCE auth.user_id_seq TO finance_calc;
