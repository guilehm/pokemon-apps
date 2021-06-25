import { PokemonDetailResult } from '../services/types'
import Badge from './Badge'
import Link from 'next/link'

import * as S from './PokemonCard.styles'

const PokemonCard: React.FC<PokemonDetailResult> = ({ id, name, sprites, types }) => (
  <S.Section>
    <article>
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
      <S.List>
        {types.map(type =>
          <Badge
            key={`${name}-type-${type.type.name}`}
            text={type.type.name}
            pokemonType={type.type.name}
          />)}
      </S.List>
    </article>
  </S.Section>
)


export default PokemonCard
