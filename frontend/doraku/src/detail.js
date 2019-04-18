import React, { Component } from 'react';
import { connect } from 'react-redux';
import { HobbyHandler } from './hobbyhandler';

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
        console.log(res)
        this.setState({
          hobby: res
        })
      }
    })
  }

  render() {
    return (
      <div>
        <h2>詳細</h2>
        <h2>ERROR: {this.state.error}</h2>
      </div>
    )
  }
}

const mapStateToProps = state => ({
  error: state.error
})

export default connect(
  mapStateToProps
)(Detail)