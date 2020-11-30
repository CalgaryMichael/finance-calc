interface IScenario {
  startDate: string
}

export type State = {
  //scenario: IScenario
  testValue: string
}

export type Action = {
  type: string
}

export type DispatchType = (args: Action) => Action

