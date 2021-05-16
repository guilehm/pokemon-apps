from fastapi import FastAPI
from fastapi import Request
from fastapi.responses import JSONResponse

app = FastAPI()


@app.get('/')
def index(request: Request):
    return JSONResponse({
      'success': True,
    })

