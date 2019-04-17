import React, { Component } from 'react';
import { connect } from 'react-redux';
import { setError } from './actions';
import { HobbyHandler } from './hobbyhandler';

class List extends Component {
  state = {
    hobbies: []
  };

  constructor(props) {
    super(props)

    let handler = new HobbyHandler()
    handler.getAllHobby().then(res => {
      if (!res) {
        let error = handler.getError()
        this.props.setError(error)
        this.props.history.push('/error')
      } else {
        this.setState({ hobbies: res })
      }
    })
  }
  render() {
    return (
      <div>
        <h2>趣味一覧</h2>
        <ul>
          {this.state.hobbies.map(item =>
            <li key={item.id}>{item.name}</li>
          )}
        </ul>
      </div>
    )
  }
}

const mapStateToProps = state => ({
  hobby: state.hobby
})

const mapDispatchToProps = {
  setError
}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(List)
