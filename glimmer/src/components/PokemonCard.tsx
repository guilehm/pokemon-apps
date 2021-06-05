import { PokemonDetailResult } from '../services/types'
import Link from 'next/link'

import * as S from './PokemonCard.styles'

const PokemonCard: React.FC<PokemonDetailResult> = ({ id, name, sprites }) => (
  <S.Section>
    <header>
      <figure>
        <S.Wrap>
          {sprites.other['official-artwork'].front_default && <S.PokemonImage
            src={sprites.other['official-artwork'].front_default}
            layout={'fill'} />}
        </S.Wrap>
        <figcaption>
          <Link href={`/pokemon/${id}/`}>
            <a><S.Name>{name}</S.Name></a>
          </Link>
        </figcaption>
      </figure>
    </header>
  </S.Section>
)


export default PokemonCard
