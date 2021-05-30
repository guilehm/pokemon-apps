import { useEffect, useState } from 'react'
import ApiService from './services/api-service'
import { PokemonListResult } from './services/types'

const Api = new ApiService()


const App = () => {

  const [pokemonList, setPokemonList] = useState<PokemonListResult[] | []>([])

  useEffect(() => {
    const fetchPokemonList = async () => {
      const response = await Api.getPokemonList()
      setPokemonList(response.data.results)
    }
    fetchPokemonList()
  }, [])


  return (
    <div className="App">
      <h1>Hello World</h1>
      {pokemonList.map((result: PokemonListResult) =>
        <a key={result.name} href={`${result.url}`}><h2>{result.name}</h2></a>
      )}

    </div>
  )
}

export default App
