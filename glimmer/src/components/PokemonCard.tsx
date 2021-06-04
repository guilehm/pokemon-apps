import { PokemonDetailResult } from '../services/types'
import Image from 'next/image'

import * as S from './PokemonCard.styles'

const PokemonCard: React.FC<PokemonDetailResult> = ({ name, sprites }) => (
  <S.Section>
    <S.Name>{name}</S.Name>
    <S.Wrap>
      <S.PokemonImage
        src={sprites.other['official-artwork'].front_default}
        layout={'fill'} />
    </S.Wrap>
  </S.Section>
)


export default PokemonCard
