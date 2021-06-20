import asyncio

import httpx
from fastapi import FastAPI, HTTPException
from fastapi.responses import JSONResponse

from routers import pokemon
from settings import PROXY_URL

app = FastAPI()
app.include_router(prefix='/fastapp/api', router=pokemon.api_router)
app.include_router(prefix='/fastapp', router=pokemon.router)
