import asyncio
import json

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
                aio_pika.Message(
                    body=json.dumps(message).encode() if isinstance(message, dict) else message.encode()
                ),
                routing_key=routing_key,
            ) for message in messages])
        return results
