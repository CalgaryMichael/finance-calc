import React from "react";
import { useSelector, useDispatch } from "react-redux"
import { Dispatch } from "redux"

import HomePage from "./views/home-page";
import { State } from "./types.d";
import * as Actions from "./store/action-creators";

const App: React.FC = (props) => {
  const projections: readonly object[] = useSelector((state: State) => state.projections);

  const dispatch: Dispatch<any> = useDispatch();

  const updateScenario = React.useCallback(
    (scenario: any) => dispatch(Actions.updateScenario(scenario)),
    [dispatch]
  );

  const saveScenario = React.useCallback(
    () => dispatch(Actions.saveScenario()),
    [dispatch]
  );

  return (
    <HomePage
      projections={projections}
      updateScenario={updateScenario}
      saveScenario={saveScenario}
    />
  );
};

export default App;
