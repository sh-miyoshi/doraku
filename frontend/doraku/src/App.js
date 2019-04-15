import React, { Component } from 'react';
import { BrowserRouter, Link, Route } from "react-router-dom";
import Top from './top'
import { Today } from './today'
import { createBrowserHistory } from 'history'
import { applyMiddleware, compose, createStore } from 'redux'
import { routerMiddleware, ConnectedRouter } from 'connected-react-router'
import rootReducer from './reducer'
import { Provider } from 'react-redux'
import Error from './error'

const history = createBrowserHistory()
const store = createStore(
  rootReducer(history), // new root reducer with router state
  {},
  compose(
    applyMiddleware(
      routerMiddleware(history), // for dispatching history actions
      // ... other middlewares ...
    ),
  ),
)

class App extends Component {
  render() {
    return (
      <Provider store={store}>
        <ConnectedRouter history={history}>
          <BrowserRouter>
            <div>
              <header>
                <h3><Link to="/">LOGO</Link></h3>
              </header>

              <Route exact path="/" component={Top} />
              <Route path="/today" component={Today} />
              <Route path="/error" component={Error} />
            </div>
          </BrowserRouter>
        </ConnectedRouter>
      </Provider>
    );
  }
}

export default App;
