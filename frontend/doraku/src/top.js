import React, { Component } from 'react';
import { withRouter } from "react-router"
import { Button } from 'react-bootstrap'
import axios from 'axios'

class Top extends Component {
  handleToTodayPage = async () => {
    try {
      const response = await axios.get('http://localhost:8080/api/v1/hobby/today');
      console.log(response);
    } catch (error) {
      console.error(error)
    }
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