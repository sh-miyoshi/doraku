import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import { connect } from 'react-redux';

class User extends Component {
  render() {
    return (
      <div>
        <Link to="/">新規登録</Link>
        <Link to="/login">ログイン</Link>
      </div>
    )
  }
}

const mapStateToProps = state => ({})

export default connect(
  mapStateToProps
)(User)