import React, { Component } from 'react';
import { connect } from 'react-redux'

class Today extends Component {
  render() {
    return (
      <div>
        <h2>Today</h2>
        {this.props.hobby.selected_hobby_id}<br />
        {this.props.hobby.selected_hobby_name}<br />
      </div>
    )
  }
}

const mapStateToProps = state => ({
  hobby: state.hobby
})

export default connect(
  mapStateToProps,
)(Today)
