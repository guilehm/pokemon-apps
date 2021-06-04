import os

import motor.motor_asyncio

MONGODB_URI = os.getenv('MONGODB_URI')

DB_CLIENT = motor.motor_asyncio.AsyncIOMotorClient(MONGODB_URI)
DB_NAME = MONGODB_URI.rsplit('/')[-1]
db = DB_CLIENT[DB_NAME]
