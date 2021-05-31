import axios, { AxiosInstance } from 'axios'

const API_URL = process.env.API_URL || 'http://spoon-proxy'
const POKEMON_API_BASE_URL = process.env.REACT_APP_POKEMON_API_BASE_URL || 'fastapp/pokemon'


class ApiService {
  baseUrl: string
  client: AxiosInstance

  constructor(baseUrl: string = POKEMON_API_BASE_URL) {
    this.baseUrl = `${API_URL}/${baseUrl}`
    this.client = axios.create()
  }

  getPokemonList(limit: number = 20, offset: number = 0) {
    return this.client.get(`${this.baseUrl}/`, { params: { limit, offset } })
  }

}


export default ApiService
