import React, { Component } from 'react';
import { connect } from 'react-redux';
import { Button } from 'react-bootstrap';
// import { Link } from 'react-router-dom';
// import { setInternalServerError } from './actions';
// import { HobbyHandler } from './hobbyhandler';

class Recommend extends Component {
  render() {
    return (
      <div>
        <h1>質問</h1>
        <ul>
          <li>Q1: 外？</li>
          <li>Q2: 一人で？</li>
          <li>Q3: アクティブ？</li>
        </ul>
        <Button onClick={() => { this.props.history.push('/recommend_result') }}>
          診断
        </Button>
      </div>
    )
  }
}

const mapStateToProps = state => ({
  error: state.error
})

// const mapDispatchToProps = {
//   setInternalServerError
// }

export default connect(
  mapStateToProps,
  //  mapDispatchToProps
)(Recommend)
