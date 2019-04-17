import React, { Component } from 'react';
import { connect } from 'react-redux';
//import { HobbyHandler } from './hobbyhandler';

class Detail extends Component {
  render() {
    return (
      <div>
        <h2>詳細</h2>
        {this.props.match.params.id}
      </div>
    )
  }
}

const mapStateToProps = state => ({
  error: state.error
})

export default connect(
  mapStateToProps
)(Detail)