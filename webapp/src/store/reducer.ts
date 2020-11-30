import * as ActionType from "./action-types";
import { Action, State } from "../types.d";

const initialState: State = {
  scenario: ''
};

const reducer = (
  state: State = initialState,
  action: Action
): State => {
  switch (action.type) {
    case ActionType.UPDATE_SCENARIO:
      return {...state, ...action.payload};
    default:
      console.log(`Unable to reduce action ${action.type}`);
      return state;
  }
};

export default reducer;

