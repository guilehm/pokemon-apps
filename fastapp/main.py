import asyncio
from random import uniform, randint

import httpx
from fastapi import FastAPI, HTTPException
from fastapi.responses import JSONResponse

API_URL = 'http://spoon-proxy'

app = FastAPI()


@app.get('/')
async def index():
    async def make_get_request(url, _index):
        async with httpx.AsyncClient(timeout=30) as client:
            response = await client.get(url, params={'time': _index})
            try:
                response.raise_for_status()
            except (httpx.RequestError, httpx.HTTPStatusError):
                raise HTTPException(status_code=400)
        return response.json()

    urls = (f'{API_URL}/api/delay/' for _ in range(10))
    results = await asyncio.gather(*[make_get_request(url, i) for i, url in enumerate(urls, 1)])
    return JSONResponse(results)
