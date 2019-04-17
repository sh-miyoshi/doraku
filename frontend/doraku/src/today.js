import React, { Component } from 'react';
import { connect } from 'react-redux';
import { Link } from 'react-router-dom';
import { setInternalServerError } from './actions';
import { HobbyHandler } from './hobbyhandler';

class Today extends Component {
  state = {
    hobby_id: 0,
    hobby_name: ""
  };

  constructor(props) {
    super(props)

    let handler = new HobbyHandler()
    handler.getTodayHobby().then(res => {
      if (!res) {
        let error = handler.getError()
        this.props.setInternalServerError(error)
        this.props.history.push('/error')
      } else {
        console.log(res)
        this.setState({
          hobby_id: res.id,
          hobby_name: res.name
        })
      }
    })
  }

  render() {
    return (
      <div>
        <h2>今日の趣味はこれ！</h2>
        <h1>
          <Link to={this._getPath()}>
            {this.state.hobby_name}
          </Link>
        </h1>
      </div>
    )
  }

  _getPath = () => {
    return "/detail/" + this.state.hobby_id
  }
}

const mapStateToProps = state => ({
  error: state.error
})

const mapDispatchToProps = {
  setInternalServerError
}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Today)
