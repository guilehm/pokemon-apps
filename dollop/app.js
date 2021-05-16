const express = require('express')
const morgan = require('morgan')
const helmet = require('helmet')
const logger = require('./utils/logger')

const PORT = process.env.PORT || 3000
const DEBUG = process.env.DEBUG || 1


const app = new express()
app.use(express.urlencoded({ extended: true }))
app.use(morgan('short'))
app.use(helmet())


app.post('/api/delay/', (req, res) => {
  const { time } = req.body
  setTimeout(() => res.json({ time: `${time}s` }), time * 1000)
})

app.listen(PORT, () => {
  const message = parseInt(DEBUG) ? 'Starting development server on port' : 'App listening on port'
  logger.info(`${message} ${PORT}`)
})
