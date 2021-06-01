import express from 'express'
import helmet from 'helmet'
import logger, { winstonMiddleware } from './utils/logger'


const PORT = process.env.PORT || 3000

const app = express()

app.use(express.urlencoded({ extended: true }))
app.use(winstonMiddleware)
app.use(helmet())

app.listen(PORT, () => {
  logger.info(`App listening on port ${PORT}`)
})
