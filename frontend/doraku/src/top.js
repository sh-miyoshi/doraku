import React, { Component } from 'react';
import { Button } from 'react-bootstrap'
import axios from 'axios'
import { connect } from 'react-redux'
import { setHobby, setError } from './actions'

class Top extends Component {
  handleToTodayPage = async () => {
    let ok = false
    try {
      let response = await axios.get('http://localhost:8080/api/v1/hobby/today');
      console.log(response);
      if (response && response.status === 200) {
        ok = true
        // TODO(set response.data.id, name to store)
      } else {
        this.props.setError(response)
      }
    } catch (error) {
      console.error(error)
      this.props.setError("failed to get hobby from a server")
    }
    if (ok) {
      this.props.history.push('/today')
    } else {
      console.log(this.props)
      this.props.history.push('/error')
    }
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
  hobby: state.hobby
})

const mapDispatchToProps = {
  setHobby,
  setError
}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Top)
