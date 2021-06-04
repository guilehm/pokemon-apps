import { PokemonDetailResult } from "../services/types"
import Image from 'next/image'

const PokemonCard: React.FC<PokemonDetailResult> = ({ name, sprites }) => (
  <div key={name}>
    <h2>{name}</h2>
    <Image
      src={sprites.other["official-artwork"].front_default}
      width={100}
      height={100} />
  </div>
)


export default PokemonCard
