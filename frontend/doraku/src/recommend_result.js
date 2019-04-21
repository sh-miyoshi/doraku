import React, { Component } from 'react';
import { connect } from 'react-redux';
import { Link } from 'react-router-dom';
// import { setInternalServerError } from './actions';
// import { HobbyHandler } from './hobbyhandler';

class RecommendResult extends Component {
  render() {
    return (
      <div>
        <h2>今あなたにオススメの趣味はこれ！</h2>
        <h1>
          <Link to={this._getPath()}>
            hobby
          </Link>
        </h1>
      </div>
    )
  }

  _getPath = () => {
    let id = "1"
    return "/detail/" + id
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
)(RecommendResult)
