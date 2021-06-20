import styled, { css } from 'styled-components'
import { BadgeType, PokemonType } from '../types/Pokemon';


const typePallet = {
  bug: '#44bd32',
  dark: '#2f3640',
  dragon: '#192a56',
  electric: '#fbc531',
  fairy: '#f8a5c2',
  fighting: '#e15f41',
  fire: '#e84118',
  flying: '#786fa6',
  ghost: '#8c7ae6',
  grass: '#4cd137',
  ground: '#e1b12c',
  ice: '#0097e6',
  normal: '#dcdde1',
  poison: '#9c88ff',
  psychic: '#c44569',
  rock: '#353b48',
  steel: '#7f8fa6',
  water: '#00a8ff',
}


export const Badge = styled.span<BadgeType>`
 display: inline-block;
  font-size: .75rem;
  font-weight: 700;
  color: white;
  background-color: ${props => props.pokemonType ? typePallet[props.pokemonType] : 'gray'};
  padding: .35em .65em;
  text-align: center;
  border-radius: .25rem;
  white-space: nowrap;
  line-height: 1;
  margin: 2px;
`;
