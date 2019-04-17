import React, { Component } from 'react';
import { BrowserRouter, Link, Route, Switch } from 'react-router-dom';
import { createBrowserHistory } from 'history';
import { applyMiddleware, compose, createStore } from 'redux';
import { routerMiddleware, ConnectedRouter } from 'connected-react-router';
import { Provider } from 'react-redux';
import rootReducer from './reducer';
import Error500 from './errors/500';
import { Top } from './top';
import Today from './today';
import { Detail } from './detail';
import List from './list';
import { Error404 } from './errors/404'

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

              <Switch>
                <Route exact path="/" component={Top} />
                <Route path="/today" component={Today} />
                <Route path="/error" component={Error500} />
                <Route path="/detail/:id" component={Detail} />
                <Route path="/list" component={List} />
                <Route component={Error404} />
              </Switch>
            </div>
          </BrowserRouter>
        </ConnectedRouter>
      </Provider>
    );
  }
}

export default App;
