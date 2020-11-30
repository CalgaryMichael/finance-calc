import React from "react";
import { useSelector, useDispatch } from "react-redux"
import { Dispatch } from "redux"

import HomePage from "./views/home-page";
import { State } from "./types.d";
import { runTest } from "./store/action-creators";

const App: React.FC = (props) => {
  const testValue: string = useSelector((state: State) => state.testValue);

  const dispatch: Dispatch<any> = useDispatch();

  const test = React.useCallback(
    () => dispatch(runTest()),
    [dispatch]
  );

  return <HomePage test={test} testValue={testValue} />;
};

export default App;

