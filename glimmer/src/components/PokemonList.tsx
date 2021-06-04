import { PokemonDetailResult } from '../services/types'
import PokemonCard from './PokemonCard'
import * as S from './PokemonList.styles'


interface PokemonDetailResultProps {
  pokemonList: PokemonDetailResult[]
}

const PokemonList: React.FC<PokemonDetailResultProps> = ({ pokemonList }) => (
  <S.PokemonList>
    {pokemonList.map(pokemon =>
      <S.PokemonListItem key={pokemon.name}>
        <PokemonCard {...pokemon} />
      </S.PokemonListItem>)}
  </S.PokemonList>
)


export default PokemonList
