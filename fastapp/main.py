import asyncio
import os

import httpx
from fastapi import FastAPI, HTTPException
from fastapi.responses import JSONResponse

from routers import pokemon
from settings import PROXY_URL

app = FastAPI()
app.include_router(prefix='/fastapp', router=pokemon.router)


@app.get('/')
async def index():
    async def make_get_request(url, _index):
        async with httpx.AsyncClient() as client:
            response = await client.get(url, params={'time': _index})
            try:
                response.raise_for_status()
            except (httpx.RequestError, httpx.HTTPStatusError):
                raise HTTPException(status_code=400)
        return response.json()

    urls = (f'{PROXY_URL}/api/delay/' for _ in range(10))
    results = await asyncio.gather(
        *[make_get_request(url, idx) for idx, url in enumerate(urls, 1)]
    )
    return JSONResponse(results)
