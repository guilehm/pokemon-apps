import styled from 'styled-components';


export const PokemonList = styled.ul`
display: grid;
grid-template-columns: repeat(auto-fit, minmax(160px, 1fr));

@media (min-width: 576px) {
  grid-template-columns: repeat(auto-fit, minmax(180px, 1fr));
}

`;


export const PokemonListItem = styled.li`
box-shadow: 0px 0px 40px -25px rgba(0,0,0,0.75);
margin: 5px;
`;
