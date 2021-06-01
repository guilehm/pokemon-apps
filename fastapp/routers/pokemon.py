import asyncio

import httpx
from fastapi import APIRouter

from utils import make_request

API_BASE_URL = 'https://pokeapi.co/api/v2'

router = APIRouter(
    prefix='/pokemon',
    responses={404: {'description': 'Not found'}},
)


@router.get('/')
async def pokemon_list(limit: int = 20, offset: int = 0):
    async with httpx.AsyncClient() as client:
        response = await make_request(
            client=client,
            url=f'{API_BASE_URL}/pokemon/',
            params={'limit': limit, 'offset': offset},
        )
        pokemon_list_data = response.json()['results']
        results = await asyncio.gather(
            *[make_request(client, data['url']) for data in pokemon_list_data]
        )
    return [result.json() for result in results]
