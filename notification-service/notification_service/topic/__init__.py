import json
from asyncio import AbstractEventLoop

from jsonschema import validate

from aiokafka import AIOKafkaConsumer

from notification_service.dot_env import DotEnv
import notification_service.service.email_service as email_service


class Topics:
    ORDER_CONFIRMATION_EMAIL = 'order-confirmation-email'


order_confirmation_schema = {
    "type": "object",
    "properties": {
        "id": {"type": "string"},
        "user": {
            "type": "object",
            "properties": {
                "email": {"type": "email"},
                "name": {"type": "string"},
                "address": "#/$defs/address"
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
                "email": {"type": "email"},
                "address": "#/$defs/address"
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
                "number": {"type": "string"},
            }
        }
    }
}


def value_deserializer(v):
    return json.loads(v).encode('utf-8')


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
                print(
                    "{}:{:d}:{:d}: key={} value={} timestamp_ms={}".format(
                        msg.topic, msg.partition, msg.offset, msg.key, msg.value,
                        msg.timestamp)
                )
                # email_service.send_order_confirmation_email(msg.value)
                # validate(instance=msg, schema=order_confirmation_schema)
        except json.JSONDecodeError:
            print("Erro ao decodificar a mensagem JSON")
        finally:
            await consumer.stop()
