import { DebtPayment } from "./debt-paymnet";

export interface Debt {
  name: string;
  total: float;
  payments: Array<DebtPayment>;
  interestRate: float;
}

