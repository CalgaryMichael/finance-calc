import { SavingsPayment } from "./savings-payment";

export interface SavingsAccount {
  name: string;
  apy: float;
  initialCapital: float;
  payments: Array<SavingsPayment>;
  projectedDate: string;
}

