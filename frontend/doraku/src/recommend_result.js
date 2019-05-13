import React, { Component } from 'react';
import { connect } from 'react-redux';
import { Link } from 'react-router-dom';

class RecommendResult extends Component {
  render() {
    return (
      <div>
        <br />
        <br />
        <br />
        <br />
        <br />
        <h2>今あなたにオススメの趣味はこれ！</h2>
        <Link to={this._getPath()} className="hobby_link">
          <h1 className="hobby_name">
            {this.props.recommend.hobby_name}
          </h1>
        </Link>
        <Link to="/">戻る</Link>
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
