import * as ActionType from "./action-types";
import * as external from "../external";
import { DispatchType, Projection, Scenario, State } from "../types.d";

export function updateScenario(scenario: any) {
  return (dispatch: DispatchType) => {
    dispatch({
      type: ActionType.UPDATE_SCENARIO,
      payload: { scenario: scenario as Scenario }
    });
  };
}

export function saveScenario() {
  return (dispatch: DispatchType, getState: () => State) => {
    const { scenario } = getState();
    external.projectScenario(scenario, "Payments", true)
      .then((ret) => {
        dispatch({
          type: ActionType.SAVE_PROJECTIONS,
          payload: {
            projections: ret.data.projections as Array<Projection>
          }
        });
      })
      .catch(() => console.log("not cool"));
  };
}

