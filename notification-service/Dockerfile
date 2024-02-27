FROM python:3.11-alpine

RUN mkdir notification_service
WORKDIR /notification_service

RUN apk add --no-cache bash gcc build-base zlib-dev

RUN python3 -m pip install --upgrade pip
RUN python3 -m pip install --user pipx
ENV PATH="${PATH}:/root/.local/bin"

RUN python3 -m pipx install poetry

COPY . .
RUN poetry install

EXPOSE 8000