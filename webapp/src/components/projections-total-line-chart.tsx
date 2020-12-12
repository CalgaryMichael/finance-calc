import React from "react";

import { CartesianGrid, Line, LineChart, Tooltip, XAxis, YAxis } from "recharts";

import * as utils from "../utils";


type Props = {
  projections: Array<Projection>
}


const ProjectionsTotalLineChart: React.FC = (props: Props) => {
  const data = props.projections.map((projection) => {
    const debtReducer = (acc, current) => current.debtTotal + acc;
    const savingsReducer = (acc, current) => current.savingsTotal + acc;
    return {
      date: utils.shortDate(projection.effectiveDate),
      debtSum: projection.debtProjections.reduce(debtReducer, 0),
      savingsSum: projection.savingsProjections.reduce(savingsReducer, 0)
    };
  });

  return (
    <div>
      <h2>Totals</h2>
      <LineChart width={1000} height={400} data={data}>
        <Line type="monotone" dataKey="debtSum" stroke="#8884d8" dot={false} />
        <Line type="monotone" dataKey="savingsSum" stroke="#82ca9d" dot={false} />
        <XAxis dataKey="date" minTickGap={30} />
        <YAxis />
        <CartesianGrid stroke="#ccc" />
        <Tooltip
          formatter={(value, name, props) => {
            const formattedCurrency = Number(value).toLocaleString("en-US", { maximumFractionDigits: 2 });
            return [`$${formattedCurrency}`, name];
          }}
        />
      </LineChart>
    </div>
  );
};

export default ProjectionsTotalLineChart;

