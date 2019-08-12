import React, { Component } from 'react';
import { Button } from 'react-bootstrap';
import './top.css';
import top_logo from '../assets/images/top_logo.png';
import ReactTooltip from 'react-tooltip'

export class Top extends Component {
  render() {
    return (
      <div>
        <img src={top_logo} alt="logo" width={400} />
        <p>さぁ、新しい人生を切り開こう！</p>
        <p>Dorakuは皆様の日常に新しい変化を付け加えるため、あなたにあった趣味を提案します</p>
        <p>日々の生活に新しい変化を求めているなら、ぜひこのサイトで新しい趣味を探してみてください！</p>
        <br />
        <table cellPadding={5} align="center">
          <tbody>
            <tr>
              <td>
                <Button className="top_button" data-tip="いくつかの質問からあなたにあった趣味を提案いたします" onClick={() => { this.props.history.push('/recommend') }}>
                  趣味診断
              </Button>
              </td>
              <td>
                <Button className="top_button" data-tip="今人気の趣味を紹介します" onClick={() => { this.props.history.push('/today') }}>
                  今日の趣味
              </Button>
              </td>
              <td>
                <Button className="top_button" data-tip="Dorakuで管理している趣味の一覧を表示します" onClick={() => { this.props.history.push('/list') }}>
                  趣味一覧
              </Button>
              </td>
            </tr>
          </tbody>
        </table>
        <ReactTooltip effect="solid" place="bottom" />
      </div>
    )
  }
}
