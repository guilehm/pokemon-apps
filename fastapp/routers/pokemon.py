import asyncio

import httpx
from fastapi import APIRouter

from services.publisher import RabbitMQService
from settings import POKEMON_API_BASE_URL, POKEMON_ROUTING_KEY
from utils import make_request

router = APIRouter(
    prefix='/pokemon',
    responses={404: {'description': 'Not found'}},
)


@router.get('/')
async def pokemon_list(limit: int = 20, offset: int = 0):
    async with httpx.AsyncClient() as client:
        response = await make_request(
            client=client,
            url=f'{POKEMON_API_BASE_URL}/pokemon/',
            params={'limit': limit, 'offset': offset},
        )
        service = RabbitMQService()
        results = await asyncio.gather(
            *[make_request(client, result['url']) for result in response.json()['results']],
        )
        results_data = [result.json() for result in results]

    await service.publish_messages(
        POKEMON_ROUTING_KEY,
        (data for data in results_data),
    )
    return results_data
