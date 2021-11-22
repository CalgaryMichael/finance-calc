CREATE TABLE finance.debt_payment (
	id		SERIAL PRIMARY KEY,
	debt_id		BIGINT REFERENCES finance.debt(id),
	amount		NUMERIC(14,2) NOT NULL,
	carry_over	BOOLEAN DEFAULT FALSE,
	start_date	DATE,
	end_date	DATE
);

GRANT ALL ON TABLE finance.debt_payment TO finance_calc;
GRANT ALL ON SEQUENCE finance.debt_payment_id_seq TO finance_calc;
