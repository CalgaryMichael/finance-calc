//interface IScenario {
//  startDate: string
//}

export type State = {
  scenario: any
}

export type Action = {
  type: string,
  payload: object
}

export type DispatchType = (args: Action) => Action

