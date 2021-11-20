import { Debt } from "./debt";
import { SavingsAccount } from "./savings-account";
import { SortKey } from "../sort-keys";

export interface Scenario {
  startDate?: string;
  debts: Array<Debt>;
  savingsAccounts: Array<SavingsAccount>;
  sortKey: string;
  reverse: boolean;
}

