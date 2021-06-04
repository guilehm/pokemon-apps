import { useState } from 'react'
import { PokemonDetailResult } from '../services/types'
import PokemonCard from './PokemonCard'
import * as S from './PokemonList.styles'


interface PokemonDetailResultProps {
  list: PokemonDetailResult[]
}

const PokemonList: React.FC<PokemonDetailResultProps> = ({ list }) => {
  const [pokemonList, setPokemonList] = useState<PokemonDetailResult[]>(list)
  return (
    <S.PokemonList>
      {pokemonList.map(pokemon =>
        <S.PokemonListItem key={pokemon.name}>
          <PokemonCard {...pokemon} />
        </S.PokemonListItem>)}
    </S.PokemonList>
  )
}


export default PokemonList
