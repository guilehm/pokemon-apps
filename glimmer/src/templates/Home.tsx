import { PokemonDetailResult } from '../services/types'
import PokemonList from '../components/PokemonList'
import * as S from './Home.styles';


interface PokemonDetailResultProps {
  pokemonList: PokemonDetailResult[]
}

const Home: React.FC<PokemonDetailResultProps> = ({ pokemonList }) => (
  <S.Container>
    <PokemonList list={pokemonList} />
  </S.Container>
)


export default Home
