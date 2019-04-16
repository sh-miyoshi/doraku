import React, { Component } from 'react';
import { Button } from 'react-bootstrap'
import { connect } from 'react-redux'
import { setHobby, setError } from './actions'
import { HobbyHandler } from './hobbyhandler'

class Top extends Component {
  handleToTodayPage = async () => {
    let handler = new HobbyHandler()
    let res = await handler.getTodayHobby()
    if (!res) {
      let error = handler.getError()
      this.props.setError(error)
      this.props.history.push('/error')
    } else {
      console.log(res)
      this.props.setHobby(res.id, res.name)
      this.props.history.push('/today')
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
