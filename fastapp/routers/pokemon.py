import asyncio
from typing import List

import httpx
from fastapi import APIRouter

from database.models import PokemonModel
from database.mongodb import db
from services.publisher import RabbitMQService
from settings import POKEMON_API_BASE_URL, POKEMON_ROUTING_KEY
from utils import make_request, make_pokemon_detail_response

api_router = APIRouter(
    prefix='/pokemon',
    responses={404: {'description': 'Not found'}},
)

router = APIRouter(
    prefix='/pokemon',
    responses={404: {'description': 'Not found'}},
)


@api_router.get('/')
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

    *results_data, _ = await asyncio.gather(
        *[make_pokemon_detail_response(data.json()) for data in results],
        service.publish_messages(
            POKEMON_ROUTING_KEY,
            (data.json() for data in results),
        )
    )
    return results_data


@router.get('/{pokemon_id}/')
async def pokemon_detail(pokemon_id: int):
    async with httpx.AsyncClient() as client:
        response = await make_request(
            client=client,
            url=f'{POKEMON_API_BASE_URL}/pokemon/{pokemon_id}',
        )
    return await make_pokemon_detail_response(response.json())


@router.get(
    '/',
    response_description='Retrieve all Pokemon from database',
    response_model=List[PokemonModel],
)
async def pokemon_list():
    pokemon_list_data = await db['pokemons'].find(
        {},
        {'id': 1, 'name': 1},
    ).to_list(2000)
    return pokemon_list_data
