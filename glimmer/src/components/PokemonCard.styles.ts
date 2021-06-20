import styled from 'styled-components';
import Image, { ImageProps } from 'next/image'


export const Section = styled.section`
  padding: 5px;
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
