import mongoose, { Schema } from 'mongoose'

type PokemonSchemaType = {
  id: number
  name: string
  order?: number
  height?: number
  weight?: number
}

const PokemonSchema: Schema = new mongoose.Schema({
  id: {
    type: Number,
    required: true,
    unique: true,
    index: true,
  },
  name: {
    type: String,
    required: true,
    unique: true,
    index: true,
  },
  order: {
    type: Number,
  },
  height: {
    type: Number,
  },
  weight: {
    type: Number,
  },
  date_added: {
    type: Date,
    default: Date.now,
  },
}, { strict: false })

const Pokemon = mongoose.model('Pokemon', PokemonSchema)


export default Pokemon
export { PokemonSchemaType }
