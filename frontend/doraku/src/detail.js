import React, { Component } from 'react';
import { connect } from 'react-redux';
import { HobbyHandler } from './hobbyhandler';
import Iframe from 'react-iframe';
import { BACKEND_SERVER_URL } from './env.secret';

class Detail extends Component {
  state = {
    hobby: null,
    error: ""
  };

  constructor(props) {
    super(props)

    let handler = new HobbyHandler()
    handler.getHobbyDetail(this.props.match.params.id).then(res => {
      if (!res) {
        let error = handler.getError()
        console.log(error)
        this.setState({
          error: error
        })
      } else {
        console.log(res.data)
        this.setState({
          hobby: res.data
        })
      }
    })
  }

  render() {
    return (
      <div>
        <h2>詳細</h2>
        {
          this.state.error &&
          <h2>ERROR: {this.state.error}</h2>
        }
        {
          this.state.hobby &&
          <h2>{this.state.hobby.name}</h2>
        }
        <img src={this._getImagePath(this.props.match.params.id)} alt="hobby" height="150" weight="150" />
        {
          this.state.hobby && this.state.hobby.description ?
            <Iframe url={this.state.hobby.description}
              width="450px"
              height="450px"
              id="myId"
              className="myClassname"
              display="initial"
              position="relative" />
            : <h2>説明文はありません</h2>
        }
      </div>
    )
  }

  _getImagePath = (id) => {
    return BACKEND_SERVER_URL + "/api/v1/hobby/image/" + id
  }
}

const mapStateToProps = state => ({
  error: state.error
})

export default connect(
  mapStateToProps
)(Detail)