import { SavingsPayment } from "./savings-payment";

export type SavingsAccount = {
  name: string,
  apy: float,
  initialCapital: float,
  payments: Array<SavingsPayment>,
  projectedDate: string
}

