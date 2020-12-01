import axios from "axios";

import * as ActionType from "./action-types";
import * as Settings from "../settings";
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
    const payload = {
      scenario: JSON.parse(scenario),
      sortKey: "Payments",
      reverse: true,
    };
    axios.post(`${Settings.API_BASE}/project`, payload)
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

