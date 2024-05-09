import os
from dotenv import load_dotenv

load_dotenv()

class AppConfig:
    COUCHDB_USER = os.getenv("COUCHDB_USER", "admin")
    COUCHDB_PWD = os.getenv("COUCHDB_PWD", "123456")
    COUCHDB_HOST = os.getenv("COUCHDB_HOST", "127.0.0.1")
    COUCHDB_PORT: int = os.getenv("COUCHDB_PORT", 5984)  # type: ignore
    POSTGRES_URI: str = os.getenv("POSTGRES_URI", "")
    COUCHDB_URI = f"http://{COUCHDB_USER}:{COUCHDB_PWD}@{COUCHDB_HOST}:{COUCHDB_PORT}"
    COUCHDB_NAME = os.getenv("COUCHDB_NAME")
    POSTGRES_NAMESPACE = os.getenv("POSTGRES_NAMESPACE")
    PERSIST_DIR = os.getenv("PERSIST_DIR")
    EMBED_MODEL = os.getenv("EMBED_MODEL")
    EMBED_MODEL_PATH = os.getenv("EMBED_MODEL_PATH")
    EMBED_MODEL_NAME = os.getenv("EMBED_MODEL_NAME")
    EMBED_MODEL_DIM = os.getenv("EMBED_MODEL_DIM")
    TOGETHER_API_KEY = os.getenv("TOGETHER_API_KEY")
