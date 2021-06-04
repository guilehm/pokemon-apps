import amqp from 'amqplib/callback_api'
import logger from './utils/logger'
import PokemonService from './services/pokemon-service'
import mongo from './database/mongodb'

const RABBITMQ_URL = process.env.RABBITMQ_URL
const POKEMON_ROUTING_KEY = process.env.POKEMON_ROUTING_KEY


amqp.connect(RABBITMQ_URL, (connError, connection) => {

  if (connError) throw connError

  mongo

  connection.createChannel((error, channel) => {
    if (error) throw error
    const pokemonService = new PokemonService()
    channel.assertQueue(POKEMON_ROUTING_KEY, { durable: false })
    logger.info(` [*] Waiting for messages in ${POKEMON_ROUTING_KEY}.`,)
    channel.consume(POKEMON_ROUTING_KEY, (msg) => {
      if (!msg) return
      logger.info(` [x] Received ${msg && msg.properties.messageId}`)
      pokemonService.updateOrCreatePokemon(JSON.parse(msg.content.toString()))
    }, { noAck: true })
  })
})
