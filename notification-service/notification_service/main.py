import asyncio

import uvicorn
from aiokafka import AIOKafkaConsumer
from fastapi import FastAPI
from dotenv import load_dotenv

from notification_service.topic import Topic

load_dotenv()

app = FastAPI()

loop = asyncio.get_event_loop()


@app.on_event("startup")
def startup_app():
    asyncio.create_task(Topic.send_order_confirmation_email_consumer(loop))


@app.get("/")
def home():
    return "Hello World!"


def run():
    uvicorn.run("notification_service.main:app", reload=True)
