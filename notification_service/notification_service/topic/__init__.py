from asyncio import AbstractEventLoop

import jsons
from aiokafka import AIOKafkaConsumer
from jsonschema.validators import validate

from notification_service.dot_env import DotEnv
from notification_service.model.order import Order
from notification_service.service import email_service


class Topics:
    ORDER_CONFIRMATION_EMAIL = 'order-confirmation-email'


order_confirmation_schema = {
    "type": "object",
    "properties": {
        "id": {"type": "string"},
        "user": {
            "type": "object",
            "properties": {
                "email": {"type": "string", "format": "email"},
                "name": {"type": "string"},
                "address": {"$ref": "#/$defs/address"}
            }
        },
        "products": {
            type: "array",
            "items": {
                "type": "object",
                "properties": {
                    "name": {"type": "string"},
                    "price": {"type": "number"},
                    "quantity": {"type": "integer"}
                }
            }
        },
        "company": {
            "type": "object",
            "properties": {
                "name": {"type": "string"},
                "email": {"type": "string", "format": "email"},
                "address": {"$ref": "#/$defs/address"}
            }
        }
    },
    "$defs": {
        "address": {
            "type": "object",
            "properties": {
                "street": {"type": "string"},
                "city": {"type": "string"},
                "state": {"type": "string"},
                "number": {"type": "integer"},
            }
        }
    }
}


def value_deserializer(v):
    data = jsons.loadb(v)
    validate(instance=data, schema=order_confirmation_schema)
    return jsons.load(data, Order)


class Topic:
    @staticmethod
    async def send_order_confirmation_email_consumer(loop: AbstractEventLoop):
        consumer = AIOKafkaConsumer(
            Topics.ORDER_CONFIRMATION_EMAIL,
            loop=loop,
            bootstrap_servers=DotEnv.kafka_server(),
            value_deserializer=value_deserializer
        )

        await consumer.start()

        try:
            async for msg in consumer:
                email_service.send_order_confirmation_email(msg.value)
        except jsons.JsonsError:
            print("Error decoding JSON message")
        finally:
            await consumer.stop()
