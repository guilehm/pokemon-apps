import mongo, { Mongoose, UpdateWriteOpResult } from 'mongoose'
import Pokemon, { PokemonSchemaType } from '../database/models/pokemon-model'
import logger from '../utils/logger'

class PokemonService {

  db: Mongoose

  constructor() {
    this.db = mongo
  }

  async updateOrCreatePokemon(pokemonData: PokemonSchemaType): Promise<UpdateWriteOpResult> {
    logger.info(`Updating Pokemon ${JSON.stringify(pokemonData['name'])}`)
    return Pokemon.updateOne(
      { id: pokemonData.id },
      pokemonData,
      { upsert: true, setDefaultsOnInsert: true },
    )
  }

}


export default PokemonService
