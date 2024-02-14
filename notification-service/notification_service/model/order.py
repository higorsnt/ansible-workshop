from dataclasses import dataclass

from notification_service.model.company import Company
from notification_service.model.product import Product
from notification_service.model.user import User


@dataclass
class Order:
    id: str
    user: User
    products: list[Product]
    company: Company
