import { PokemonDetailResult } from '../services/types'
import PokemonCard from '../components/PokemonCard'
import * as S from './Home.styles';

interface PokemonDetailResultProps {
  pokemonList: PokemonDetailResult[]
}


const Home: React.FC<PokemonDetailResultProps> = ({ pokemonList }) => (
  <S.Container>
    <S.PokemonList>
      {pokemonList.map(pokemon =>
        <S.PokemonListItem key={pokemon.name}>
          <PokemonCard {...pokemon} />
        </S.PokemonListItem>)}
    </S.PokemonList>
  </S.Container>
)


export default Home
