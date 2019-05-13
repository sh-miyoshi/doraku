import React, { Component } from 'react';
import { connect } from 'react-redux';
import { Button, Form } from 'react-bootstrap';
import { setRecommendHobby, setInternalServerError } from './actions';
import { HobbyHandler } from './hobbyhandler';
import "./recommend.css";

class Recommend extends Component {
  state = {
    selectValues: ["yes", "yes", "yes"],// outdoor, alone, active
    nowQuestion: 0
  }

  render() {
    return (
      <div>
        <Form>
          <br />
          <br />
          <br />
          <h1>質問</h1>

          {
            this.state.nowQuestion === 0 &&
            <div>
              Q1: どっち派？
              <Form.Check label="アウトドア派" type="radio" name="1" value="yes" defaultChecked={true} onChange={this._change} />
              <Form.Check label="インドア派" type="radio" name="1" value="no" onChange={this._change} />
              <Button onClick={this._next}>次の質問へ</Button>
            </div>
          }

          {
            this.state.nowQuestion === 1 &&
            <div>
              Q2: 気分的には？
              <Form.Check label="一人な気分" type="radio" name="2" value="yes" defaultChecked={true} onChange={this._change} />
              <Form.Check label="みんなで集まりたい気分" type="radio" name="2" value="no" onChange={this._change} />
              <Button onClick={this._back}>前の質問へ</Button>
              <Button onClick={this._next}>次の質問へ</Button>
            </div>
          }

          {
            this.state.nowQuestion === 2 &&
            <div>
              Q3: どんな感じがタイプ？
              <Form.Check label="激しい感じで" type="radio" name="3" value="yes" defaultChecked={true} onChange={this._change} />
              <Form.Check label="落ち着いた感じで" type="radio" name="3" value="no" onChange={this._change} />
              <Button onClick={this._back}>前の質問へ</Button>
              <Button className="enter_button" onClick={this._handleRecommend}>診断</Button>
            </div>
          }
        </Form>
      </div>
    )
  }

  _handleRecommend = async () => {
    console.log("select: " + this.state.selectValues)// for debug

    // set query params
    let params = {
      outdoor: this.state.selectValues[0],
      alone: this.state.selectValues[1],
      active: this.state.selectValues[2],
    }
    let handler = new HobbyHandler()
    let res = await handler.getRecommendHobby(params)

    if (!res) {
      let error = handler.getError()
      this.props.setInternalServerError(error)
      this.props.history.push('/error')
    } else {
      console.log(res.data)
      this.props.setRecommendHobby(res.data.id, res.data.name)
      this.props.history.push('/recommend_result')
    }
  }

  _change = (e) => {
    let n = parseInt(e.target.name, 10) - 1
    let values = this.state.selectValues.slice()
    if (0 <= n && n < values.length) {
      values[n] = e.target.value

      this.setState({
        selectValues: values
      })
    }
  }

  _back = () => {
    if (this.state.nowQuestion >= 1) {
      this.setState({
        nowQuestion: this.state.nowQuestion - 1
      })
    }
  }

  _next = () => {
    if (this.state.nowQuestion < 2) {
      this.setState({
        nowQuestion: this.state.nowQuestion + 1
      })
    }
  }
}

const mapStateToProps = state => ({
  error: state.error
})

const mapDispatchToProps = {
  setRecommendHobby,
  setInternalServerError
}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Recommend)
