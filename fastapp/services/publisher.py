import asyncio

import aio_pika

from settings import RABBITMQ_URL


class RabbitMQService:

    def __init__(self, url=RABBITMQ_URL):
        self.url = url

    async def publish_messages(self, routing_key, messages):
        connection = await aio_pika.connect_robust(self.url)
        async with connection:
            channel = await connection.channel()
            results = await asyncio.gather(*[channel.default_exchange.publish(
                aio_pika.Message(body=message.encode()),
                routing_key=routing_key,
            ) for message in messages])
        return results
