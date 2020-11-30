//interface IScenario {
//  startDate: string
//}

export type State = {
  scenario: any,
  projections: Array<object>
}

export type Action = {
  type: string,
  payload: object
}

export type DispatchType = (args: Action) => Action

