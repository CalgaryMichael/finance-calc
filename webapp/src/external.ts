import axios, { AxiosResponse } from "axios";

import * as Settings from "./settings";

export const projectScenario = (scenario: Scenario, sortKey: string, reverse: boolean): Promise<AxiosResponse> => {
  const payload = {
    scenario: JSON.parse(scenario),
    sortKey,
    reverse
  };
  return axios.post(`${Settings.API_BASE}/project`, payload);
}

