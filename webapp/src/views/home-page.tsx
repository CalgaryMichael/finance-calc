import React from "react";

import { Projection } from "../types.d";
import ProjectionsTotalLineChart from "../components/projections-total-line-chart";


type Props = {
  projections: Array<Projection>
  updateScenario: (scenario: any) => void,
  saveScenario: () => void
}

const HomePage: React.FC = (props: Props) => {
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
      <div>
        <textarea
          onBlur={updateScenario}
        />
      </div>
      <button onClick={saveScenario}>Submit</button>
      <div>
        <ProjectionsTotalLineChart
          projections={props.projections}
        />
        <pre>{ JSON.stringify(props.projections, null, 2) }</pre>
      </div>
    </div>
  );
};

export default HomePage;
