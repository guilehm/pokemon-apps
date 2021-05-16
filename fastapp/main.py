import asyncio
from random import randint

import httpx
from fastapi import FastAPI, HTTPException
from fastapi import Request
from fastapi.responses import JSONResponse

DOLLOP_URL = 'http://dollop'

app = FastAPI()


@app.get('/')
async def index(request: Request):
    async def make_post_request(url):
        async with httpx.AsyncClient(timeout=30) as client:
            response = await client.post(url, params={'time': randint(1, 3)})
            try:
                response.raise_for_status()
            except (httpx.RequestError, httpx.HTTPStatusError):
                raise HTTPException(status_code=400)
        return response.json()

    urls = (f'{DOLLOP_URL}/api/delay/' for _ in range(20))
    results = await asyncio.gather(*[make_post_request(url) for url in urls])
    return JSONResponse(results)
