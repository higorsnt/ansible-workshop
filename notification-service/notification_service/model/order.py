from notification_service.model.company import Company
from notification_service.model.product import Product
from notification_service.model.user import User


class Order:
    def __init__(self, id: str, user: User, products: [Product], company: Company):
        self.id = id
        self.user = user
        self.products = products
        self.company = company
