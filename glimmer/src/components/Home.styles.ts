import styled from 'styled-components';


export const Container = styled.div`
  max-width: 80%;
  margin: auto;
`;


export const PokemonList = styled.ul`
  display: flex;
  align-items: center;
  justify-content: space-between;
  flex-wrap: wrap;
`;


export const PokemonListItem = styled.li`
  list-style: none;
  width: 200px;
  height: 200px;
`;
