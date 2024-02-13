from notification_service.model.address import Address


class Company:
    def __init__(self, name: str, email: str, address: Address):
        self.name = name
        self.email = email
        self.address = address
