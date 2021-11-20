import * as ActionType from "./action-types";
import { Action, SortKeys, State } from "../types.d";

const initialState: State = {
  scenario: {
    startDate: null,
    debts: [],
    savingsAccounts: [],
    sortKey: SortKeys.Payments,
    reverse: true
  },
  projections: []
};

const reducer = (
  state: State = initialState,
  action: Action
): State => {
  switch (action.type) {
    case ActionType.UPDATE_SORT_KEY:
      return {...state, scenario: {...state.scenario, sortKey: action.payload}};
    case ActionType.UPDATE_SORT_DIRECTION:
      return {...state, scenario: {...state.scenario, reverse: action.payload}};
    case ActionType.UPDATE_SCENARIO:
      return {
        ...state,
        scenario: {
          startDate: null,
          debts: [],
          savingsAccounts: [],
          ...action.payload,
          sortKey: state.scenario.sortKey,
          reverse: state.scenario.reverse
        }
      };
    case ActionType.SAVE_PROJECTIONS:
      return {...state, projections: action.payload.projections || []};
    default:
      console.log(`Unable to reduce action ${action.type}`);
      return state;
  }
};

export default reducer;

