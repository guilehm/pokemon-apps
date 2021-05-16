const Express = require('express')
const morgan = require('morgan')
const helmet = require('helmet')
const logger = require('./utils/logger')


const app = new Express()

app.use(morgan('short'))
app.use(helmet())

const PORT = process.env.PORT || 3000
const DEBUG = process.env.DEBUG || 1


app.get('/', (req, res) => res.json({ ok: true }))

app.listen(PORT, () => {
  const message = parseInt(DEBUG) ? 'Starting development server on port' : 'App listening on port'
  logger.info(`${message} ${PORT}`)
})
