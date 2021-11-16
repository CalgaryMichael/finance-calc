CREATE TABLE finance.savings_payment (
	id			SERIAL PRIMARY KEY,
	savings_account_id	BIGINT REFERENCES finance.savings_account(id),
	amount			NUMERIC(14,2) NOT NULL DEFAULT 0,
	start_date		DATE,
	end_date		DATE
);

GRANT ALL ON TABLE finance.savings_payment TO finance_calc;
GRANT ALL ON SEQUENCE finance.savings_payment_id_seq TO finance_calc;
