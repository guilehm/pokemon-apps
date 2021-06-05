type PokemonListResult = {
  id: string,
  name: string,
}

type PokemonDetailResult = {
  name: string
  sprites: {
    other: {
      'official-artwork': {
        'front_default': string
      }
    }
  }
}


export type { PokemonDetailResult, PokemonListResult }
