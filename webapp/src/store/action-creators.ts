import axios from "axios";

import * as ActionType from "./action-types";
import * as Settings from "../settings";
import { DispatchType, State } from "../types.d";

export function updateScenario(scenario: any) {
  return (dispatch: DispatchType) => {
    dispatch({
      type: ActionType.UPDATE_SCENARIO,
      payload: { scenario }
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
      .then(() => console.log("cool"))
      .catch(() => console.log("not cool"));
  };
}

