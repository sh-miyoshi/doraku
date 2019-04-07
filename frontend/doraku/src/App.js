import React, { Component } from 'react';
import { BrowserRouter as Router, Route } from "react-router-dom";
import { Top } from './top'

class App extends Component {
  render() {
    return (
      <Router>
        <div>
          <header>
            <h3>LOGO</h3>
          </header>

          <Route exact path="/" component={Top} />
        </div>
      </Router>
    );
  }
}

export default App;
