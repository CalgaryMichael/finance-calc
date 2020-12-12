import * as ActionType from "./action-types";
import * as external from "../external";
import { buildProjections, DispatchType, Scenario, State } from "../types.d";

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
            projections: buildProjections(ret.data.projections)
          }
        });
      })
      .catch((e) => console.log(e));
  };
}

