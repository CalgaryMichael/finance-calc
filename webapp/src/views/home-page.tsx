import React from "react";

import { Projection, Scenario, SortKeys } from "../types.d";
import ProjectionsTotalLineChart from "../components/projections-total-line-chart";


type Props = {
  scenario: Scenario,
  projections: Array<Projection>,
  updateSortKey: (sortKey: string) => void,
  updateSortDirection: (reverse: boolean) => void,
  updateScenario: (scenario: any) => void,
  saveScenario: () => void
}

const HomePage: React.FC = (props: Props) => {
  const updateSortKey = (e: React.ChangeEvent<HTMLInputElement>) => {
    e.preventDefault();
    props.updateSortKey(e.target.value);
  };
  const updateSortDirection = (e: React.ChangeEvent<HTMLInputElement>) => {
    e.preventDefault();
    props.updateSortDirection(e.target.value === "true");
  };
  const updateScenario = (e: React.ChangeEvent<HTMLInputElement>) => {
    e.preventDefault();
    props.updateScenario(e.target.value);
  };
  const saveScenario = (e: React.FormEvent) => {
    e.preventDefault();
    props.saveScenario();
  };

  return (
    <div>
      <div style={{display: "flex"}}>
        <select name="sortKey" value={props.scenario.sortKey} onChange={updateSortKey}>
          <option value={SortKeys.DebtName}>Debt Name</option>
          <option value={SortKeys.Payments}>Payments</option>
          <option value={SortKeys.DebtTotal}>Debt Total</option>
          <option value={SortKeys.InterestRate}>Interest Rate</option>
        </select>

        <input
          type="radio"
          id="sort-direction-asc"
          name="sort-direction"
          value={false}
          checked={props.scenario.reverse === false}
          onChange={updateSortDirection}
        />
        <label htmlFor="sort-direction-asc">ASC</label>

        <input
          type="radio"
          id="sort-direction-desc"
          name="sort-direction"
          value={true}
          checked={props.scenario.reverse === true}
          onChange={updateSortDirection}
        />
        <label htmlFor="sort-direction-desc">DESC</label>
      </div>
      <div>
        <textarea onBlur={updateScenario} />
      </div>
      <button onClick={saveScenario}>Submit</button>
      <ProjectionsTotalLineChart projections={props.projections} />
    </div>
  );
};

export default HomePage;

