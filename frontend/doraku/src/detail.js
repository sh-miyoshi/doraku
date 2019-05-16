import React, { Component } from 'react';
import { connect } from 'react-redux';
import { HobbyHandler } from './hobbyhandler';
import { BACKEND_SERVER_URL } from './env.secret';
import './detail.css'

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
            <p>{this.state.hobby.description}</p>
            : <h2>説明文は準備中です・・・</h2>
        }
        {
          this.state.hobby && this.state.hobby.descriptionFrom &&
          <p>from {this.state.hobby.descriptionFrom}</p>
        }
        {
          this.state.hobby && this.state.hobby.descriptionURL &&
          <p>詳細は<a href={this.state.hobby.descriptionURL}>こちら</a></p>
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