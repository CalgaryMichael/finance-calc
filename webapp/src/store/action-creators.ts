import * as ActionType from "./action-types";
import * as external from "../external";
import { buildProjections, DispatchType, Scenario, State } from "../types.d";

export function updateSortKey(sortKey: string) {
  return (dispatch: DispatchType) => {
    dispatch({
      type: ActionType.UPDATE_SORT_KEY,
      payload: sortKey
    });
  }
}

export function updateSortDirection(reverse: boolean) {
  return (dispatch: DispatchType) => {
    dispatch({
      type: ActionType.UPDATE_SORT_DIRECTION,
      payload: reverse
    });
  }
}

export function updateScenario(scenario: string) {
  return (dispatch: DispatchType) => {
    dispatch({
      type: ActionType.UPDATE_SCENARIO,
      payload: (scenario ? JSON.parse(scenario) : {}) as Scenario
    });
  };
}

export function saveScenario() {
  return (dispatch: DispatchType, getState: () => State) => {
    const { scenario } = getState();
    external.projectScenario(scenario)
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

