import styled from 'styled-components';
import Image, { ImageProps } from 'next/image'


export const Section = styled.section`
`;

export const Name = styled.h2`
  text-align: center;
`;

export const Wrap = styled.div`
  position: relative;
  width: 150px;
  height: 150px;
  margin: auto;
`;

export const PokemonImage = styled(Image) <ImageProps>`
`;
