import os

PROXY_URL = os.getenv('API_URL', 'http://pokemon.local.com')
RABBITMQ_URL = os.getenv('RABBITMQ_URL')
POKEMON_API_BASE_URL = 'https://pokeapi.co/api/v2'

POKEMON_ROUTING_KEY = os.getenv('POKEMON_ROUTING_KEY')
