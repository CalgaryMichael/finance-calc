CREATE TABLE finance.debt_projection (
	debt_id		BIGINT REFERENCES finance.debt(id),
	effective_date	DATE NOT NULL,
	total		NUMERIC(14,2) NOT NULL DEFAULT 0,
	payment_sum	NUMERIC(14,2) NOT NULL DEFAULT 0,
	unapplied_sum	NUMERIC(14,2) NOT NULL DEFAULT 0,

	PRIMARY KEY (debt_id, effective_date)
);

GRANT ALL ON TABLE finance.debt_projection TO finance_calc;
