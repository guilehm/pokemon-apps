# Pokemon

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

  ```
  127.0.0.1 pokemon.local.com
  ```


Start the apps

    docker-compose up

Wait for all services to come up and you may access the app at this url: [http://pokemon.local.com](http://pokemon.local.com/)



## Services

### dollop

An Node.js worker to listen to a RabbitMQ queue.

Receive messages and create or update documents in MongoDB.

Made with TypeScript.

---
### fastapp

Python with FastAPI backend running with Uvicorn.

Retrieve data from [PokéAPI](https://pokeapi.co/) and publish responses in RabbitMQ queue.

Also have endpoints that retrieve data from MongoDB.

---
### goapp

An GO alternative backend that makes the same of "fastapp".

You can choose which backend to use.
Just change the environment variable `API_URL` in "glimmer" at file `api-service.tsx`.

GO is faster.

---
### glimmer

A React / Next.js frontend.

Made with TypeScript and Styled Components.

---
### proxy

A reverse proxy using NGINX.

Responsible to proxy pass and cache requests to the apps.

---

## Troubleshooting

If you have MongoDB and or RabbitMQ, running in your machine, please stop it before running the app:

    sudo systemctl stop mongodb


## Contribution
If you want to contribute, just open an issue.
Fork the repository and change whatever you'd like.

Pull requests are always welcome.
