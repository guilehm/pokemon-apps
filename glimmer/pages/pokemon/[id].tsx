import { GetStaticPaths } from 'next'
import PokemonCard from '../../src/components/PokemonCard'
import ApiService from '../../src/services/api-service'
import { PokemonDetailResult, PokemonListResult } from '../../src/services/types'


const Api = new ApiService()

const PokemonDetailPage = (context: PokemonDetailResult) => {
  return (
    <PokemonCard {...context} />
  )
}

export default PokemonDetailPage

type Params = {
  params: {
    id: number
  }
}

export const getStaticProps = async ({ params }: Params) => {
  if (!params || !params.id) return
  const response = await Api.getPokemonDetail(params.id)
  return {
    props: response.data,
    revalidate: 60,
  }
}

export const getStaticPaths: GetStaticPaths<{ id: string }> = async () => {
  const response = await Api.getPokemonList()
  return {
    paths: response.data.map((result: PokemonListResult) =>
      ({ params: { id: result.id.toString() } })),
    fallback: 'blocking'
  }
}
