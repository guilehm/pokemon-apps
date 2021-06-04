import React from 'react'
import PokemonCard from '../src/components/PokemonCard'
import ApiService from '../src/services/api-service'
import { PokemonDetailResult } from '../src/services/types'


interface PokemonDetailResultProps {
  pokemonList: PokemonDetailResult[]
}

const Home: React.FC<PokemonDetailResultProps> = ({ pokemonList }) => {
  return (
    <div>
      <main>
        <h1>
          Hello World!
        </h1>
        {pokemonList.map(pokemon => <PokemonCard key={pokemon.name} {...pokemon} />)}
      </main>
    </div>
  )
}


export default Home


export const getStaticProps = async () => {
  const Api = new ApiService()
  const response = await Api.getPokemonList()
  const pokemonList = response.data || []
  return {
    props: {
      pokemonList
    },
    revalidate: 60,
  }
}
