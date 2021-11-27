import axios, { AxiosResponse } from "axios";

import * as Settings from "./settings";

export const saveScenario = (scenario: Scenario): Promise<AxiosResponse> => {
  const payload = {
    scenario
  };
  return axios.post(`${Settings.API_BASE}/scenario`, payload);
}

export const getProjections = (scenarioId: number): Promise<AxiosResponse> => {
  return axios.get(`${Settings.API_BASE}/scenario/${scenarioId}/projections`);
}
