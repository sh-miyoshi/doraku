import axios from 'axios';

const SERVER_URL = "http://localhost:8080"

export class HobbyHandler {
  constructor() {
    this.error = null
  }

  async getTodayHobby() {
    return await this._query(SERVER_URL + "/api/v1/hobby/today")
  }

  async getAllHobby() {
    return await this._query(SERVER_URL + "/api/v1/hobby/all")
  }

  async getHobbyDetail(id) {
    return await this._query(SERVER_URL + "/api/v1/hobby/details/" + id)
  }

  async getRecommendHobby(params) {
    return await this._query(SERVER_URL + "/api/v1/hobby/recommended", params)
  }

  getError() {
    return this.error
  }

  // private methods
  async _query(url, params = null) {
    try {
      console.log(params)
      let response = await axios.get(url, { params: params });
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