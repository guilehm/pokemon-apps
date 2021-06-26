import { BadgeType } from '../types/Pokemon'

import * as S from './Badge.styles'


const Badge: React.FC<BadgeType> = ({ text, color, pokemonType }) => (
  <S.Badge color={color} pokemonType={pokemonType}>{text}</S.Badge>
)


export default Badge
