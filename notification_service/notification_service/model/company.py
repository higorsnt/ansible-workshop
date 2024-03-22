from dataclasses import dataclass

from notification_service.model.address import Address


@dataclass
class Company:
    name: str
    email: str
    address: Address
