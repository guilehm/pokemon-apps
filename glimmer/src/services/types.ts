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
      name: string,
    }
  }[]
}


export type { PokemonDetailResult, PokemonListResult }
