FROM python:3.11-alpine

RUN mkdir notification_service
WORKDIR /notification_service

COPY /pyproject.toml /notification_service
RUN pip3 install pipx
RUN pipx install poetry
RUN poetry install

COPY . .

EXPOSE 8000