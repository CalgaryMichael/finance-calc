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
    const { reverse, scenario, sortKey } = getState();
    external.projectScenario(scenario, sortKey, reverse)
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

