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
    errorMessage: '',
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
                this.setState({
                  name: e.target.value
                })
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
          <div className="error">{this.state.errorMessage}</div>
          <Button variant="primary" type="submit">ログイン</Button>
        </Form>
      </div>
    )
  }

  _handleLogin = (e) => {
    console.log("login with " + this.state.name + " : " + this.state.password)

    if (!this._validateName(this.state.name)) {
      this.setState({
        errorMessage: "正しいユーザ名を入力してください"
      })
      e.preventDefault()
      return
    }
  }

  _validateName = (name) => {
    if (!name) {// If name is null or undefined
      return false
    }
    const val = "" + name

    const pattern = /[^a-zA-Z0-9.\-_]/
    if (val.match(pattern)) {
      return false
    }

    return (4 <= val.length && val.length < 32)
  }
}

const mapStateToProps = state => ({})

const mapDispatchToProps = {}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Login)
