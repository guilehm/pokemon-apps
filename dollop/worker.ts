import amqp from 'amqplib/callback_api'
import logger from './utils/logger'


const RABBITMQ_URL = process.env.RABBITMQ_URL!
const POKEMON_ROUTING_KEY = process.env.POKEMON_ROUTING_KEY!


amqp.connect(RABBITMQ_URL, (connError, connection) => {

  if (connError) throw connError

  connection.createChannel((error, channel) => {
    if (error) throw error

    channel.assertQueue(POKEMON_ROUTING_KEY, { durable: false })

    logger.info(` [*] Waiting for messages in ${POKEMON_ROUTING_KEY}.`,)

    channel.consume(POKEMON_ROUTING_KEY, (msg) => {
      logger.info(` [x] Received ${msg!.properties.messageId}`)
    }, { noAck: true })

  })
})

