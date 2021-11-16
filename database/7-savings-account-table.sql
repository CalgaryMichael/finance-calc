CREATE TABLE finance.savings_account (
	id		SERIAL PRIMARY KEY,
	name		TEXT,
	apy		NUMERIC(2,2) NOT NULL DEFAULT 0.0,
	initial_capital	NUMERIC(14,2) NOT NULL DEFAULT 0,
	projected_date	DATE
);

GRANT ALL ON TABLE finance.savings_account TO finance_calc;
GRANT ALL ON SEQUENCE finance.savings_account_id_seq TO finance_calc;
