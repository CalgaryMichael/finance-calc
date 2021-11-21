import axios, { AxiosResponse } from "axios";

import * as Settings from "./settings";

export const projectScenario = (scenario: Scenario): Promise<AxiosResponse> => {
  const payload = {
    scenario
  };
  return axios.post(`${Settings.API_BASE}/scenario`, payload);
}
