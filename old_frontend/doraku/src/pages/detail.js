import React, { Component } from 'react';
import { connect } from 'react-redux';
import { HobbyHandler } from '../plugins/hobbyhandler';
import { BACKEND_SERVER_URL } from '../plugins/global_constant';
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
        <br />
        <br />
        <br />
        <br />
        {
          this.state.error &&
          <h2 className="error" >ERROR: {this.state.error}</h2>
        }
        {
          this.state.hobby &&
          <h2>{this.state.hobby.name}</h2>
        }
        <img src={this._getImagePath(this.props.match.params.id)} alt="hobby" height="150" weight="150" />
        <div className="description_box">
          {
            this.state.hobby && this.state.hobby.description ?
              <p>{this.state.hobby.description}</p>
              : <p>説明文は準備中です・・・</p>
          }
        </div>
        {
          this.state.hobby && this.state.hobby.descriptionFrom &&
          <span className="where">from {this.state.hobby.descriptionFrom}</span>
        }
        {
          this.state.hobby && this.state.hobby.descriptionURL &&
          <div className="descriptionURL">詳細は<a href={this.state.hobby.descriptionURL}>こちら</a></div>
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