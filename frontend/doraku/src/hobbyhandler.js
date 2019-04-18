import axios from 'axios';

const SERVER_URL = "http://localhost:8080"

export class HobbyHandler {
  constructor() {
    this.error = null
  }

  async getTodayHobby() {
    let res = await this._query(SERVER_URL + "/api/v1/hobby/today")
    if (res) {
      return { id: res.data.id, name: res.data.name }
    }
    return null
  }

  async getAllHobby() {
    let res = await this._query(SERVER_URL + "/api/v1/hobby/all")
    if (res) {
      return res.data
    }
    return null
  }

  async getHobbyDetail(id) {
    let res = await this._query(SERVER_URL + "/api/v1/hobby/details/" + id)
    if (res) {
      return res.data
    }
    return null
  }

  getError() {
    return this.error
  }

  // private methods
  async _query(url) {
    try {
      let response = await axios.get(url);
      console.log(response);
      if (response && response.status === 200) {
        return response
      } else {
        this.error = response
      }
    } catch (error) {
      console.error(error)
      this.error = "failed to get hobby from a server"
    }
    return null
  }
}