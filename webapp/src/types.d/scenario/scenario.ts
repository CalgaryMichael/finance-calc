import { Debt } from "./debt";
import { SavingsAccount } from "./savings-account";

export interface Scenario {
  startDate: string;
  debts: Array<Debt>;
  savingsAccounts: Array<SavingsAccount>;
}

