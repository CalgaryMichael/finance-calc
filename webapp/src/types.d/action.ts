import { Projection } from "./projection";
import { Scenario } from "./scenario";

export type ScenarioPayload = {
  scenario: Scenario
}

export type ProjectionPayload = {
  projection: Array<Projection>
}

export type Action = {
  type: string,
  payload: ProjectionPayload | ScenarioPayload
}

