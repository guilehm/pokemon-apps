import mongoose from 'mongoose'
import logger from '../utils/logger'


const MONGODB_URI = process.env.MONGODB_URI as string
mongoose.connect(MONGODB_URI, {
  useNewUrlParser: true,
  useFindAndModify: false,
  useCreateIndex: true,
  useUnifiedTopology: true,
})
  .then(() => logger.info('Connected to Mongo'))
  .catch(e => logger.error('Error while trying to connect to Mongo', e))


export default mongoose
