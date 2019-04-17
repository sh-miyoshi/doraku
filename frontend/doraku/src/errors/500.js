import React, { Component } from 'react';
import { connect } from 'react-redux';
import { Link } from 'react-router-dom';
import './500.css';

class Error500 extends Component {
  render() {
    return (
      <div className="main">
        <h1>500</h1>
        <h2>Unexpected Error <b>:(</b></h2>
        <h2>{this.props.error.errorMsg}</h2>
        <h2><Link to="/">もう一度試す</Link></h2>
        <div className="gears">
          <div className="gear one">
            <div className="bar"></div>
            <div className="bar"></div>
            <div className="bar"></div>
          </div>
          <div className="gear two">
            <div className="bar"></div>
            <div className="bar"></div>
            <div className="bar"></div>
          </div>
          <div className="gear three">
            <div className="bar"></div>
            <div className="bar"></div>
            <div className="bar"></div>
          </div>
        </div>
      </div>
    )
  }
}

const mapStateToProps = state => ({
  error: state.error
})

export default connect(
  mapStateToProps,
)(Error500)
