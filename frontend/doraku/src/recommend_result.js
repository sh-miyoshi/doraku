import React, { Component } from 'react';
import { connect } from 'react-redux';
import { Link } from 'react-router-dom';

class RecommendResult extends Component {
  render() {
    return (
      <div>
        <h2>今あなたにオススメの趣味はこれ！</h2>
        <h1>
          <Link to={this._getPath()}>
            {this.props.recommend.hobby_name}
          </Link>
        </h1>
      </div>
    )
  }

  _getPath = () => {
    return "/detail/" + this.props.recommend.hobby_id
  }
}

const mapStateToProps = state => ({
  recommend: state.recommend
})

export default connect(
  mapStateToProps,
)(RecommendResult)
