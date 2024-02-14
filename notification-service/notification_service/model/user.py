from dataclasses import dataclass

from notification_service.model.address import Address


@dataclass
class User:
    email: str
    name: str
    address: Address
