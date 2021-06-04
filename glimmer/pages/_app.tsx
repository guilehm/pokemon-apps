import { AppProps } from 'next/app';
import { createGlobalStyle } from 'styled-components'
import '../styles/globals.css'


const GlobalStyle = createGlobalStyle`
  body {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
  }

  ul {
    padding: 0;
    list-style-type: none;
  }

  figure {
    margin: 0;
  }
`

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <>
      <GlobalStyle />
      <Component {...pageProps} />
    </>
  )
}

export default MyApp
