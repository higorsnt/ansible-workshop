import asyncio
import subprocess
import sys

import uvicorn
from dotenv import load_dotenv
from fastapi import FastAPI
from starlette.staticfiles import StaticFiles

from notification_service.topic import Topic

load_dotenv()

app = FastAPI()
app.mount("/async-api-doc", StaticFiles(directory="docs", html=True), name="asyncapidoc")

loop = asyncio.get_event_loop()


@app.on_event("startup")
def startup_app():
    asyncio.create_task(Topic.send_order_confirmation_email_consumer(loop))


@app.get("/")
def home():
    return "Hello World!"


def run():
    uvicorn.run("notification_service.main:app", reload=True, host="0.0.0.0")


def install_asyncapi():
    """
    Install the asyncapi module, necessary to generate documentation.
    If an error occurs regarding permissions, try this solution:
    https://stackoverflow.com/questions/18088372/how-to-npm-install-global-not-as-root
    """
    print("AsyncAPI is not installed. Installing AsyncAPI...")
    result = subprocess.run(["npm", "install", "-g", "@asyncapi/cli"])
    if result.returncode != 0:
        print("Error: Failed to install AsyncAPI.")
        sys.exit(1)


def asyncapi_doc_generator():
    try:
        subprocess.run(["asyncapi", "--version"], stdout=subprocess.PIPE, stderr=subprocess.PIPE, check=True)
    except subprocess.CalledProcessError:
        install_asyncapi()

    result = subprocess.run(["asyncapi", "generate", "fromTemplate",
                             "./notification_service/api_specification/specification.yaml", "@asyncapi/html-template",
                             "-o", "./docs"])
    if result.returncode != 0:
        print("Error: Failed to generate documentation using AsyncAPI.")
        sys.exit(1)
    sys.exit(0)
