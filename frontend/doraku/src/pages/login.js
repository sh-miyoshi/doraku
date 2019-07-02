import React, { Component } from 'react';
import { connect } from 'react-redux';
import "./login.css";
import { Button, Form } from 'react-bootstrap';
// import { setInternalServerError } from '../store/actions';
// import { HobbyHandler } from '../plugins/hobbyhandler';

class Login extends Component {
  state = {
    name: '',
    password: '',
  }

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
            <Form.Control
              type="name"
              placeholder="user name"
              value={this.state.name}
              onChange={(e) => {
                this.setState({ name: e.target.value })
              }}
            />
          </Form.Group>

          <Form.Group controlId="formBasicPassword">
            <Form.Label>パスワード</Form.Label>
            <Form.Control
              type="password"
              placeholder="Password"
              value={this.state.password}
              onChange={(e) => {
                this.setState({ password: e.target.value })
              }}
            />
          </Form.Group>
          <Button variant="primary" type="submit">ログイン</Button>
        </Form>
      </div>
    )
  }

  _handleLogin = (e) => {
    e.preventDefault();

    console.log("login with " + this.state.name + " : " + this.state.password)
  }
}

const mapStateToProps = state => ({})

const mapDispatchToProps = {}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Login)
