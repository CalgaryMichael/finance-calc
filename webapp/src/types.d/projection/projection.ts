import { LocalDate } from "js-joda";

import { parseLocalDate } from "../../utils";

import { DebtProjection } from "./debt-projection";
import { SavingsProjection } from "./savings-projection";

export class Projection {
  effectiveDate: LocalDate;
  debtProjections: Array<DebtProjection>;
  savingsProjections: Array<SavingsProjection>;

  constructor(input) {
    Object.assign(this, input);
  }

  static of(input: object): Projection {
    if (!input) {
      return null;
    }
    const { effectiveDate, ...rest } = input;
    return new Projection({
      ...rest,
      effectiveDate: parseLocalDate(effectiveDate)
    });
  }
}

export const buildProjections = (inputs: Array<object>): Array<Projection> => {
  return inputs.map((input) => Projection.of(input));
};

