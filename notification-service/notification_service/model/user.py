from notification_service.model.address import Address


class User:
    def __init__(self, email: str, name: str, address: Address):
        self.email = email
        self.name = name
        self.address = address
