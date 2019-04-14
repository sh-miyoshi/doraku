import React, { Component } from 'react';
import { connect } from 'react-redux'

class Error extends Component {
  render() {
    return (
      <div>
        <h2>Internal Server Error</h2>
        {this.props.hobby.error}
      </div>
    )
  }
}

const mapStateToProps = state => ({
  hobby: state.hobby
})

export default connect(
  mapStateToProps,
)(Error)
