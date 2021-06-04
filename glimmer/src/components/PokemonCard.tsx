import { PokemonDetailResult } from '../services/types'

import * as S from './PokemonCard.styles'

const PokemonCard: React.FC<PokemonDetailResult> = ({ name, sprites }) => (
  <S.Section>
    <header>
      <figure>
        <S.Wrap>
          {sprites.other['official-artwork'].front_default && <S.PokemonImage
            src={sprites.other['official-artwork'].front_default}
            layout={'fill'} />}
        </S.Wrap>
        <figcaption>
          <S.Name>{name}</S.Name>
        </figcaption>
      </figure>
    </header>
  </S.Section>
)


export default PokemonCard
