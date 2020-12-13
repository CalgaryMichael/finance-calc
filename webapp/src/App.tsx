import React from "react";
import { useSelector, useDispatch } from "react-redux"
import { Dispatch } from "redux"

import HomePage from "./views/home-page";
import { State } from "./types.d";
import * as Actions from "./store/action-creators";

const App: React.FC = (props) => {
  const sortKey: string = useSelector((state: State) => state.sortKey);
  const reverse: boolean = useSelector((state: State) => state.reverse);
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
      sortKey={sortKey}
      reverse={reverse}
      projections={projections}
      updateSortKey={updateSortKey}
      updateSortDirection={updateSortDirection}
      updateScenario={updateScenario}
      saveScenario={saveScenario}
    />
  );
};

export default App;
