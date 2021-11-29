CREATE TABLE finance.savings_projection (
	savings_account_id	BIGINT REFERENCES finance.savings_account(id),
	effective_date		DATE NOT NULL,
	total			NUMERIC(14,2) NOT NULL DEFAULT 0,
	payment_sum		NUMERIC(14,2) NOT NULL DEFAULT 0,

	PRIMARY KEY (savings_account_id, effective_date)
);

GRANT ALL ON TABLE finance.savings_projection TO finance_calc;
