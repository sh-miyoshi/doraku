import React, { Component } from 'react';
import { withRouter } from "react-router"
import { Button } from 'react-bootstrap'

class Top extends Component {
  handleToTodayPage = () => {
    this.props.history.push('/today')
  }

  render() {
    return (
      <div>
        <h1>LOGO</h1>
        <Button onClick={this.handleToTodayPage}>
          今日の趣味
        </Button>
      </div>
    )
  }
}

export default withRouter(Top)