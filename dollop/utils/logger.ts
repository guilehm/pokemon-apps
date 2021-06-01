import expressWinston from 'express-winston'
import winston from 'winston'


const loggerOptions = {
  format: winston.format.combine(
    winston.format.timestamp({ format: 'YYYY-MM-DD HH:mm:ss:ms' }),
    winston.format.colorize({ all: true }),
    winston.format.printf(
      (info) => `${info.timestamp} ${info.level}: ${info.message}`,
    ),
  ),
  transports: [new winston.transports.Console()]
}

const logger = winston.createLogger(loggerOptions)
const winstonMiddleware = expressWinston.logger(loggerOptions)

export default logger
export { winstonMiddleware }
