import ApiService from '../src/services/api-service'

const Home = () => {
  return (
    <div>
      <main>
        <h1>
          Hello World!
        </h1>
      </main>
    </div>
  )
}


export default Home


export const getStaticProps = async () => {
  const Api = new ApiService()
  return {
    props: {
    },
    revalidate: 60,
  }
}
