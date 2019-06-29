import React, { Component } from 'react';
import { connect } from 'react-redux';
import "./login.css";
// import { setInternalServerError } from '../store/actions';
// import { HobbyHandler } from '../plugins/hobbyhandler';

class Login extends Component {
  render() {
    return (
      <div className="itembox">
        <br />
        <br />
        <br />
        <br />
        <table>
          <tbody>
            <tr>
              <td>ユーザ名</td>
              <td><input type="text" /></td>
            </tr>
            <tr>
              <td>パスワード</td>
              <td><input type="password" /></td>
            </tr>
          </tbody>
        </table>
      </div>
    )
  }
}

const mapStateToProps = state => ({})

const mapDispatchToProps = {}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Login)
