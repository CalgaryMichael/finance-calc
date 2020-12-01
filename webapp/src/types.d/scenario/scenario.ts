import { Debt } from "./debt";
import { SavingsAccount } from "./savings-account";

export type Scenario = {
  startDate: string,
  debts: Array<Debt>,
  savingsAccounts: Array<SavingsAccount>
}

