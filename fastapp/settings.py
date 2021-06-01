import os

PROXY_URL = os.getenv('API_URL', 'http://spoon-proxy')
RABBITMQ_URL = os.getenv('RABBITMQ_URL')
POKEMON_API_BASE_URL = 'https://pokeapi.co/api/v2'
