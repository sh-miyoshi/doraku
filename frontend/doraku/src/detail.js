import React, { Component } from 'react';
import { connect } from 'react-redux'

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
  hobby: state.hobby
})

export default connect(
  mapStateToProps,
)(Detail)
