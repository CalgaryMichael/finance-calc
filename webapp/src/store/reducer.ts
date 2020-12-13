import * as ActionType from "./action-types";
import { Action, State } from "../types.d";

const initialState: State = {
  scenario: null,
  projections: [],
  sortKey: "Payments",
  reverse: true
};

const reducer = (
  state: State = initialState,
  action: Action
): State => {
  switch (action.type) {
    case ActionType.UPDATE_SORT_KEY:
      return {...state, sortKey: action.payload};
    case ActionType.UPDATE_SORT_DIRECTION:
      return {...state, reverse: action.payload};
    case ActionType.UPDATE_SCENARIO:
      return {...state, ...action.payload};
    case ActionType.SAVE_PROJECTIONS:
      return {...state, projections: action.payload.projections || []};
    default:
      console.log(`Unable to reduce action ${action.type}`);
      return state;
  }
};

export default reducer;

