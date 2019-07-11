import axios from 'axios';
import {
  BACKEND_SERVER_URL
} from './global_constant';

export class UserHandler {
  constructor() {
    this.error = null
  }

  async login(name, password) {
    let params = {
      name: name,
      password: password,
    }
    return await this._query('post', BACKEND_SERVER_URL + '/api/v1/login', params)
  }

  // private methods
  async _query(method, url, params = null) {
    try {
      let response = null
      switch (method.toLowerCase()) {
        case 'get':
          response = await axios.get(url, {
            params: params
          })
          break
        case 'post':
          response = await axios.post(url, params)
          break
        default:
          this.error = "no such method: " + method
          return null
      }
      console.log(response);
      if (response && (200 <= response.status && response.status < 300)) {
        return response
      } else {
        this.error = response
      }
    } catch (error) {
      console.error(error)
      this.error = "failed to request server"
    }
    return null
  }
}