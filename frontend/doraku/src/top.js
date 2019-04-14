import React, { Component } from 'react';
import { Button } from 'react-bootstrap'
import axios from 'axios'
import { connect } from 'react-redux'
import { setHobby } from './actions'

class Top extends Component {
  handleToTodayPage = async () => {
    try {
      let response = await axios.get('http://localhost:8080/api/v1/hobby/today');
      console.log(response);
      if (!response || response.status !== 200) {
        return
      }
      // TODO(set response.data.id, name to store)
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

const mapStateToProps = state => ({
  stores: state.stores
})

const mapDispatchToProps = {
  setHobby
}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Top)
