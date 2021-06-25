export type PokemonType =
  'bug' |
  'dark' |
  'dragon' |
  'electric' |
  'fairy' |
  'fighting' |
  'fire' |
  'flying' |
  'ghost' |
  'grass' |
  'ground' |
  'ice' |
  'normal' |
  'poison' |
  'psychic' |
  'rock' |
  'steel' |
  'water'

export type BadgeType = {
  text?: string,
  color?: string,
  pokemonType?: PokemonType,
}
