import { PokemonDetailResult } from '../services/types'
import PokemonCard from './PokemonCard'


interface PokemonDetailResultProps {
  pokemonList: PokemonDetailResult[]
}


const Home: React.FC<PokemonDetailResultProps> = ({ pokemonList }) => (
  <div>
    <main>
      <h1>
        Hello World!
      </h1>
      {pokemonList.map(pokemon => <PokemonCard key={pokemon.name} {...pokemon} />)}
    </main>
  </div>
)


export default Home
