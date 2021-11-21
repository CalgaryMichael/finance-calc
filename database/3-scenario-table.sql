CREATE TABLE finance.scenario (
	id		SERIAL PRIMARY KEY,
	user_id		BIGINT REFERENCES auth.user(id),
	name		TEXT NOT NULL,
	start_date	DATE,
	sort_key	TEXT NOT NULL DEFAULT 'DebtTotal',
	reverse_sort	BOOLEAN NOT NULL DEFAULT FALSE,
	created		TIMESTAMP NOT NULL DEFAULT now()
);

GRANT ALL ON TABLE finance.scenario TO finance_calc;
GRANT ALL ON SEQUENCE finance.scenario_id_seq TO finance_calc;
