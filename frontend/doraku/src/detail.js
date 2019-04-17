import React, { Component } from 'react';

export class Detail extends Component {
  render() {
    return (
      <div>
        <h2>詳細</h2>
        {this.props.match.params.id}
      </div>
    )
  }
}