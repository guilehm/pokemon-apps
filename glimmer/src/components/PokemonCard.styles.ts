import styled from 'styled-components';
import Image, { ImageProps } from 'next/image'


export const Section = styled.section`
`;

export const Name = styled.h2`
  text-align: center;
`;

export const Wrap = styled.div`
  position: relative;
  width: 135px;
  height: 135px;
  margin: auto;
`;

export const PokemonImage = styled(Image) <ImageProps>`
`;
