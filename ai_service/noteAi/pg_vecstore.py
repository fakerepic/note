from sqlalchemy import make_url

from llama_index.vector_stores.postgres import PGVectorStore

from noteAi.config import AppConfig

connection_string = AppConfig.POSTGRES_URI
url = make_url(connection_string)
db_name = "vector_db"

vector_store = PGVectorStore.from_params(
    database=db_name,
    host=url.host,
    password=url.password,
    port=url.port, # type: ignore
    user=url.username,
    table_name="paul_graham_essay",
    embed_dim=384,  # fastembed bge embedding dimension
)


def get_vectorstore():
    return vector_store
