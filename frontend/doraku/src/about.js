import React, { Component } from 'react';
import { connect } from 'react-redux';
import './about.css'

class About extends Component {
  render() {
    return (
      <div>
        <br />
        <br />
        <br />
        <br />
        <h1>作者の想い</h1>
        <div className="message">
          <p>世の中にはきっと面白いことがたくさんある</p>
          <p>でもそのことを知らないと面白さに気づくことすらできない</p>
          <p>Dorakuはふと暇だなと思った瞬間に新しいもの気づける場所を目指して作成しています</p>
        </div>
      </div>
    )
  }
}

const mapStateToProps = state => ({})

export default connect(
  mapStateToProps
)(About)