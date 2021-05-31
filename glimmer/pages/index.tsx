import ApiService from '../src/services/api-service'
import { PokemonListResult } from '../src/services/types'

interface PokemonListResultProps {
  pokemonList: PokemonListResult[]
}

const Home: React.FC<PokemonListResultProps> = ({ pokemonList }) => {
  return (
    <div>
      <main>
        <h1>
          Hello World!
        </h1>
        {pokemonList.map(pokemon => <div key={pokemon.name}>{pokemon.name}</div>)}
      </main>
    </div>
  )
}


export default Home


export const getStaticProps = async () => {
  const Api = new ApiService()
  const response = await Api.getPokemonList()
  const pokemonList = response.data?.results || []
  return {
    props: {
      pokemonList
    },
    revalidate: 60,
  }
}
