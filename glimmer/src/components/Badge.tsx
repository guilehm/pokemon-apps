import { BadgeType } from '../types/Pokemon'

import * as S from './Badge.styles'



const Badge: React.FC<BadgeType> = ({ text, pokemonType }) => (
  <S.Badge pokemonType={pokemonType}>{text}</S.Badge>
)


export default Badge
