import { PokemonType } from '../types/Pokemon'

type PokemonListResult = {
  id: string,
  name: string,
}

type PokemonDetailResult = {
  id: string
  name: string
  sprites: {
    other: {
      'official-artwork': {
        'front_default': string
      }
    }
  },
  types: {
    slot: number,
    type: {
      name: PokemonType,
    }
  }[]
}


export type { PokemonDetailResult, PokemonListResult }
