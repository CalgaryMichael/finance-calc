import { DebtPayment } from "./debt-paymnet";

export type Debt = {
  name: string,
  total: float,
  payments: Array<DebtPayment>,
  interestRate: float
}

