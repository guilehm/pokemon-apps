import logging
import sys
import urllib.parse

import httpx
from fastapi import HTTPException


def get_logger():
    logging.basicConfig(stream=sys.stdout, level=logging.INFO, format='%(message)s')
    return logging


logger = get_logger()


async def make_request(client, url, params=None, log=False):
    log and logger.info(
        f'Requesting {url}{f"?{urllib.parse.urlencode(params or dict())}" if params else ""}'
    )
    response = await client.get(url, params=params or dict())
    try:
        response.raise_for_status()
    except (httpx.RequestError, httpx.HTTPStatusError):
        raise HTTPException(status_code=500)
    return response


async def make_pokemon_detail_response(payload):
    return {
        'id': payload['id'],
        'name': payload['name'],
        'sprites': {
            'other': {
                'official-artwork': {
                    'front_default': payload['sprites']['other']['official-artwork']['front_default'],
                }
            }
        },
        'types': payload['types'],
    }
