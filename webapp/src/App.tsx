import React from "react";
import { useSelector, useDispatch } from "react-redux"
import { Dispatch } from "redux"

import HomePage from "./views/home-page";
import { Scenario, State } from "./types.d";
import * as Actions from "./store/action-creators";

const App: React.FC = (props) => {
  const scenario: Scenario = useSelector((state: State) => state.scenario);
  const projections: readonly object[] = useSelector((state: State) => state.projections);

  const dispatch: Dispatch<any> = useDispatch();

  const updateSortKey = React.useCallback(
    (sortKey: string) => dispatch(Actions.updateSortKey(sortKey)),
    [dispatch]
  );

  const updateSortDirection = React.useCallback(
    (reverse: boolean) => dispatch(Actions.updateSortDirection(reverse)),
    [dispatch]
  );

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
      scenario={scenario}
      projections={projections}
      updateSortKey={updateSortKey}
      updateSortDirection={updateSortDirection}
      updateScenario={updateScenario}
      saveScenario={saveScenario}
    />
  );
};

export default App;
