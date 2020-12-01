import { Projection } from "./projection";
import { Scenario } from "./scenario";

export type State = {
  scenario: Scenario,
  projections: Array<Projection>
}

