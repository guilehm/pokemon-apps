import asyncio
from random import random, randint, uniform

import httpx
from fastapi import FastAPI, HTTPException
from fastapi import Request
from fastapi.responses import JSONResponse

API_URL = 'http://spoon-proxy'

app = FastAPI()


@app.get('/')
async def index():
    async def make_get_request(url):
        async with httpx.AsyncClient(timeout=30) as client:
            response = await client.get(url, params={'time': f'{uniform(1, 3)}'})
            try:
                response.raise_for_status()
            except (httpx.RequestError, httpx.HTTPStatusError):
                raise HTTPException(status_code=400)
        return response.json()

    urls = (f'{API_URL}/api/delay/' for _ in range(20))
    results = await asyncio.gather(*[make_get_request(url) for url in urls])
    return JSONResponse(results)
