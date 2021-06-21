# Microservices with Pokemon API
![Screenshot from 2021-06-20 21-51-10](https://user-images.githubusercontent.com/33688752/122693771-acc80a00-d211-11eb-8058-012d4dc7816a.png)
## Required
* Docker
* Docker Compose

## Installation

Clone this repository

    git clone git@github.com:Guilehm/studious-spoon.git

Enter the repository

    cd studious-spoon

Add `pokemon.local.com` to your hosts

*this is necessary because of the reverse proxy*

    sudo vim /etc/hosts

Add the following line to your hosts file:

`127.0.0.1 pokemon.local.com`


Start the apps

    docker-compose up

Wait for all services to come up and you may access the app at this url: [http://pokemon.local.com](http://pokemon.local.com/)



## Services

## dollop

An Node.js worker to listen to a RabbitMQ queue.

Receive messages and create or update documents in MongoDB.

Made with TypeScript.


## fastapp

Python with FastAPI backend running with Uvicorn.

Retrieve data from [PokéAPI](https://pokeapi.co/) and publish responses in RabbitMQ queue.

Also have endpoints that retrieve data from MongoDB.


## goapp

An GO alternative backend that makes the same of "fastapp".

You can choose which backend to use.
Just change the environment variable `API_URL` in "glimmer" at file `api-service.ts`.

GO is faster.


## glimmer

A React / Next.js frontend.

Made with TypeScript and Styled Components.


## proxy

A reverse proxy using NGINX.

Responsible to proxy pass and cache requests to the apps.



## Use guide

Access the frontend at [http://pokemon.local.com](http://pokemon.local.com/)

**Retrieve pokemon list from [PokéAPI](https://pokeapi.co/):**

GO backend:
```
http://pokemon.local.com/goapp/api/pokemon/?limit=10&offset=0
```

Python backend:
```
http://pokemon.local.com/fastapp/api/pokemon/?limit=10&offset=0
```

It will make a request at [PokéAPI](https://pokeapi.co/) to get a list and then make request for the details of each Pokémon in the list, returning the list to the frontend. In the background it will create or update in MongoDB the details of each Pokemon.


You can change `limit` and `offset` parameters as you want.
But keep in mind that all requests are made at the same time. Please do not use a high value for `limit`.

Be nice to the API.

This request will be cached by NGINX.

For testing purposes, if you want to bypass the cache, you can access the apps directly by changing the domain.

GO backend:
```
http://localhost:8090/goapp/api/pokemon/?limit=10&offset=0
```

Python backend:
```
http://localhost:8000/fastapp/api/pokemon/?limit=10&offset=0
```

**Retrieve pokemon list from MongoDB:**

As mentioned before, the Pokemon results are stored in MongoDB.

If you want to access the data without hitting [PokéAPI](https://pokeapi.co/), is the same as the previous step, but removing `/api/` from the url.

GO backend:
```
http://pokemon.local.com/goapp/pokemon/?limit=10&offset=0
```

Python backend:
```
http://pokemon.local.com/fastapp/pokemon/?limit=10&offset=0
```

## Troubleshooting

If you have MongoDB and or RabbitMQ, running in your machine, please stop it before running the app:

    sudo systemctl stop mongodb


## Contribution
If you want to contribute, just open an issue.
Fork the repository and change whatever you'd like.

Pull requests are always welcome.
