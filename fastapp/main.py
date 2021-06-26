from fastapi import FastAPI

from routers import pokemon

app = FastAPI()
app.include_router(prefix='/api', router=pokemon.api_router)
app.include_router(router=pokemon.router)
