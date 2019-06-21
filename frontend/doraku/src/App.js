import React, { Component } from 'react';
import { BrowserRouter, Link, Route, Switch } from 'react-router-dom';
import { createBrowserHistory } from 'history';
import { applyMiddleware, compose, createStore } from 'redux';
import { routerMiddleware, ConnectedRouter } from 'connected-react-router';
import { Provider } from 'react-redux';
import rootReducer from './store/reducer';
import Error500 from './pages/error_500';
import { Top } from './pages/top';
import Today from './pages/today';
import Detail from './pages/detail';
import List from './pages/list';
import About from './pages/about';
import Recommend from './pages/recommend';
import RecommendResult from './pages/recommend_result';
import { Error404 } from './pages/error_404';
import './App.css';
import logo from './assets/images/logo.png';

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
                <Link to="/">
                  <img src={logo} alt="Doraku" width={150} />
                </Link>
              </header>

              <Switch>
                <Route exact path="/" component={Top} />
                <Route path="/today" component={Today} />
                <Route path="/error" component={Error500} />
                <Route path="/detail/:id" component={Detail} />
                <Route path="/list" component={List} />
                <Route path="/recommend" component={Recommend} />
                <Route path="/recommend_result" component={RecommendResult} />
                <Route path="/about" component={About} />
                <Route component={Error404} />
              </Switch>

              <footer>
                <Link to="/about">Dorakuとは？</Link>
                &emsp;&emsp;
                <a href="https://github.com/sh-miyoshi/doraku/issues/new/choose">Contact Support</a>
              </footer>
            </div>
          </BrowserRouter>
        </ConnectedRouter>
      </Provider>
    );
  }
}

export default App;
