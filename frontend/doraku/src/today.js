import React, { Component } from 'react';
import { connect } from 'react-redux';
import { Link } from 'react-router-dom';

class Today extends Component {
  render() {
    return (
      <div>
        <h2>今日の趣味はこれ！</h2>
        <h1>
          <Link to={this._getPath()}>
            {this.props.hobby.selected_hobby_name}
          </Link>
        </h1>
      </div>
    )
  }

  _getPath = () => {
    return "/detail/" + this.props.hobby.selected_hobby_id
  }
}

const mapStateToProps = state => ({
  hobby: state.hobby
})

export default connect(
  mapStateToProps,
)(Today)
