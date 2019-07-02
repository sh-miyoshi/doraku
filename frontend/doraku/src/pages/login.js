import React, { Component } from 'react';
import { connect } from 'react-redux';
import "./login.css";
import { Button, Form } from 'react-bootstrap';
// import { setInternalServerError } from '../store/actions';
// import { HobbyHandler } from '../plugins/hobbyhandler';

class Login extends Component {
  // state = {
  //   validated: false,
  // }

  render() {
    return (
      <div className="itembox">
        <br />
        <br />
        <br />
        <br />
        <h3>ログイン</h3>
        <Form onSubmit={this._handleLogin} action="/">
          <Form.Group controlId="formBasicName">
            <Form.Label>ユーザ名</Form.Label>
            <Form.Control type="name" placeholder="user name" />
          </Form.Group>

          <Form.Group controlId="formBasicPassword">
            <Form.Label>パスワード</Form.Label>
            <Form.Control type="password" placeholder="Password" />
          </Form.Group>
          <Button variant="primary" type="submit">ログイン</Button>
        </Form>
      </div>
    )
  }

  _handleLogin = (e) => {
    e.preventDefault();

    console.log("login with ...")
  }
}

const mapStateToProps = state => ({})

const mapDispatchToProps = {}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Login)
