import styled from 'styled-components';
import Image, { ImageProps } from 'next/image'


export const Section = styled.section`
  padding: 5px;
  max-width: 300px;
  margin: auto;
`;

export const Name = styled.h2`
  text-align: center;
  margin: 0;
`;

export const Wrap = styled.div`
  position: relative;
  width: 135px;
  height: 135px;
  margin: auto;
`;

export const PokemonImage = styled(Image) <ImageProps>`
`;

export const List = styled.ul`
  display: flex;
  align-items: center;
  justify-content: center;
`;
