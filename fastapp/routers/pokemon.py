import httpx
from fastapi import APIRouter
from fastapi import HTTPException

API_BASE_URL = 'https://pokeapi.co/api/v2'

router = APIRouter(
    prefix='/pokemon',
    responses={404: {'description': 'Not found'}},
)


@router.get('/')
async def pokemon_list(limit: int = 20, offset: int = 0):
    async with httpx.AsyncClient(timeout=30) as client:
        response = await client.get(
            f'{API_BASE_URL}/pokemon/',
            params={'limit': limit, 'offset': offset},
        )
        try:
            response.raise_for_status()
        except (httpx.RequestError, httpx.HTTPStatusError):
            raise HTTPException(status_code=503)
    return response.json()
