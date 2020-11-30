import * as ActionType from "./action-types";
import { Action, State } from "../types.d";

const initialState: State = {
  testValue: "off"
};

const reducer = (
  state: State = initialState,
  action: Action
): State => {
  switch (action.type) {
    case ActionType.TEST:
      return {...state, testValue: "on"};
  }

  return state;
};

export default reducer;

