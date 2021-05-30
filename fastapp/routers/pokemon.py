from fastapi import APIRouter

router = APIRouter(
    prefix='/pokemon',
    responses={404: {'description': 'Not found'}},
)


@router.get('/')
async def pokemon():
    return {'ok': True}

