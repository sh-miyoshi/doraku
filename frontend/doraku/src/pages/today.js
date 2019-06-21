import React, { Component } from 'react';
import { connect } from 'react-redux';
import { Link } from 'react-router-dom';
import { setInternalServerError } from '../store/actions';
import { HobbyHandler } from '../plugins/hobbyhandler';
import './today.css';

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
        console.log(res.data)
        this.setState({
          hobby_id: res.data.id,
          hobby_name: res.data.name
        })
      }
    })
  }

  render() {
    return (
      <div>
        <br />
        <br />
        <br />
        <br />
        <br />
        <h2>今日の趣味はこれ！</h2>
        <Link to={this._getPath()} className="hobby_link">
          <h1 className="hobby_name">
            {this.state.hobby_name}
          </h1>
        </Link>
        <Link to="/">戻る</Link>
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
