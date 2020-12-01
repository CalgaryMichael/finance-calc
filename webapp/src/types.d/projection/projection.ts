import { DebtProjection } from "./debt-projection";
import { SavingsProjection } from "./savings-projection";

export type Projection = {
  effectiveDate: string,
  debtProjections: Array<DebtProjection>,
  savingsProjections: Array<SavingsProjection>
}

