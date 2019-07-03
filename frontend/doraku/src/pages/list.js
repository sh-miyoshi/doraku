import React, { Component } from 'react';
import { connect } from 'react-redux';
import { Link } from 'react-router-dom';
import { setInternalServerError } from '../store/actions';
import { HobbyHandler } from '../plugins/hobbyhandler';
import './list.css';

const COLUMN_NUM = 2

class List extends Component {
  state = {
    hobbies: null
  };

  constructor(props) {
    super(props)

    let handler = new HobbyHandler()
    handler.getAllHobby().then(res => {
      if (!res) {
        let error = handler.getError()
        this.props.setInternalServerError(error)
        this.props.history.push('/error')
      } else {
        let t = []
        for (let i = 0; i < COLUMN_NUM; i++)
          t[i] = []
        for (let i = 0; i < res.data.length; i++) {
          t[i % COLUMN_NUM].push(res.data[i])
        }
        this.setState({ hobbies: t })
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
        <h2>趣味一覧</h2>
        <div className="itembox">
          {this.state.hobbies && this.state.hobbies.map((list, index) =>
            <div className="component" key={index}>
              <ul>
                {list.map(item =>
                  <li key={item.id}><Link to={this._getPath(item.id)}>{item.name}</Link></li>
                )}
              </ul>
            </div>
          )}
          <div className="footer"></div>
        </div>
      </div>
    )
  }

  _getPath = (id) => {
    return '/detail/' + id
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
)(List)
