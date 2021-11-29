CREATE TABLE finance.debt (
	id		SERIAL PRIMARY KEY,
	scenario_id	BIGINT REFERENCES finance.scenario(id),
	name		TEXT,
	total		NUMERIC(14,2) NOT NULL,
	interest_rate	NUMERIC(2, 2) NOT NULL DEFAULT 0,

	UNIQUE(scenario_id, name)
);

GRANT ALL ON TABLE finance.debt TO finance_calc;
GRANT ALL ON SEQUENCE finance.debt_id_seq TO finance_calc;
