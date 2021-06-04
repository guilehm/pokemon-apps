import { useEffect, useState } from 'react'
import InfiniteScroll from 'react-infinite-scroll-component'
import ApiService from '../services/api-service'
import { PokemonDetailResult } from '../services/types'
import PokemonCard from './PokemonCard'
import * as S from './PokemonList.styles'


interface PokemonDetailResultProps {
  list: PokemonDetailResult[]
}

const PokemonList: React.FC<PokemonDetailResultProps> = ({ list }) => {
  const [pokemonList, setPokemonList] = useState<PokemonDetailResult[]>(list)
  const [limit] = useState(pokemonList.length)
  const [hasMore, setHasMore] = useState(true)


  const Api = new ApiService()

  const fetchNewResults = async () => {
    const response = await Api.getApiPokemonList(
      { offset: pokemonList.length, limit }
    )
    if (!response.data.length) {
      setHasMore(false)
      return
    }
    setPokemonList([...pokemonList, ...response.data])
  }

  return (

    <InfiniteScroll
      dataLength={pokemonList.length}
      next={fetchNewResults}
      hasMore={hasMore}
      loader={<h4>Loading...</h4>}
      endMessage={
        <p style={{ textAlign: 'center' }}>
          <b>Yay! You have seen it all</b>
        </p>
      }
    >
      <S.PokemonList>
        {pokemonList.map(pokemon =>
          <S.PokemonListItem key={`${pokemon.name}`}>
            <PokemonCard {...pokemon} />
          </S.PokemonListItem>)}
      </S.PokemonList>
    </InfiniteScroll>
  )
}


export default PokemonList
