import React from "react";
import { render } from "react-dom";
import { createStore, applyMiddleware, Store } from "redux";
import { Provider } from "react-redux";
import thunk from "redux-thunk";

import App from "./App";
import { Action, DispatchType, State } from "./types.d";
import reducer from "./store/reducer";


const store: Store<State, Action> & {
  dispatch: DispatchType
} = createStore(reducer, applyMiddleware(thunk));

render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.body
);

