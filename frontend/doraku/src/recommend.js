import React, { Component } from 'react';
import { connect } from 'react-redux';
import { Button, Form } from 'react-bootstrap';
import { setRecommendHobby, setInternalServerError } from './actions';
import { HobbyHandler } from './hobbyhandler';

class Recommend extends Component {
  state = {
    selectValues: ["yes", "yes", "yes"]// outdoor, alone, active
  }

  render() {
    return (
      <div>
        <Form>
          <h1>質問</h1>
          <ul>
            <li>
              Q1: どっち派？
              <Form.Check label="アウトドア派" type="radio" name="1" value="yes" defaultChecked={true} onChange={this._change} />
              <Form.Check label="インドア派" type="radio" name="1" value="no" onChange={this._change} />
            </li>
            <li>
              Q2: 気分的には？
              <Form.Check label="一人な気分" type="radio" name="2" value="yes" defaultChecked={true} onChange={this._change} />
              <Form.Check label="みんなで集まりたい気分" type="radio" name="2" value="no" onChange={this._change} />
            </li>
            <li>
              Q3: どんな感じがタイプ？
              <Form.Check label="激しい感じで" type="radio" name="3" value="yes" defaultChecked={true} onChange={this._change} />
              <Form.Check label="落ち着いた感じで" type="radio" name="3" value="no" onChange={this._change} />
            </li>
          </ul>
          <Button onClick={this._handleRecommend}>
            診断
        </Button>
        </Form>
      </div>
    )
  }

  _handleRecommend = async () => {
    console.log("select: " + this.state.selectValues)// for debug

    let params = {}// todo
    let handler = new HobbyHandler()
    let res = await handler.getRecommendHobby(params)

    if (!res) {
      let error = handler.getError()
      this.props.setInternalServerError(error)
      this.props.history.push('/error')
    } else {
      console.log(res)
      // todo
      // this.setState({
      //   hobby_id: res.id,
      //   hobby_name: res.name
      // })
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
