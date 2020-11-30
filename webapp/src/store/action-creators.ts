import * as ActionType from "./action-types";
import { DispatchType } from "../types.d";

export function runTest() {
  return (dispatch: DispatchType) => {
    console.log('here');
    dispatch({ type: ActionType.TEST });
  };
}

