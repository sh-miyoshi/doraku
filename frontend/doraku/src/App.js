import React, { Component } from 'react';
import { BrowserRouter as Router, Route, Link } from "react-router-dom";
import Top from './top'
import { Today } from './today'

class App extends Component {
  render() {
    return (
      <Router>
        <div>
          <header>
            <h3><Link to="/">LOGO</Link></h3>
          </header>

          <Route exact path="/" component={Top} />
          <Route path="/today" component={Today} />
        </div>
      </Router>
    );
  }
}

export default App;
