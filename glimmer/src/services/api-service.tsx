import axios, { AxiosInstance, AxiosResponse } from 'axios'
import { PokemonDetailResult } from './types'

const API_URL = process.env.API_URL || 'http://pokemon.local.com/fastapp'

class ApiService {
  baseUrl: string
  client: AxiosInstance

  constructor(baseUrl: string = API_URL) {
    this.baseUrl = `${baseUrl}`
    this.client = axios.create()
  }

  getApiPokemonList({ limit = 20, offset = 0 }): Promise<AxiosResponse> {
    return this.client.get(`${this.baseUrl}/api/pokemon`, { params: { limit, offset } })
  }

  getPokemonList(): Promise<AxiosResponse> {
    return this.client.get(`${this.baseUrl}/pokemon`)
  }

  getPokemonDetail(id: string): Promise<AxiosResponse<PokemonDetailResult>> {
    return this.client.get(`${this.baseUrl}/pokemon/${id}/`)
  }


}


export default ApiService
