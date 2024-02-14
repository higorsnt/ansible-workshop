from dataclasses import dataclass


@dataclass
class Address:
    street: str
    city: str
    state: str
    number: int
