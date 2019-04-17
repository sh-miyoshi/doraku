import React, { Component } from 'react';
import { Button } from 'react-bootstrap';

export class Top extends Component {
  render() {
    return (
      <div>
        <h1>LOGO</h1>
        <Button onClick={() => { this.props.history.push('/today') }}>
          今日の趣味
        </Button>
        <br />
        <Button onClick={() => { this.props.history.push('/list') }}>
          趣味一覧
        </Button>
      </div>
    )
  }
}
