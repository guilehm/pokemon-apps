import { AppProps, AppContext } from 'next/app';
import '../styles/globals.css'


function MyApp({ Component, pageProps }: AppProps) {
  return <Component {...pageProps} />
}

export default MyApp
