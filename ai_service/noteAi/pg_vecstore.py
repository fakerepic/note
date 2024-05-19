from sqlalchemy import make_url

from llama_index.vector_stores.postgres import PGVectorStore

from noteAi.config import AppConfig

connection_string = AppConfig.POSTGRES_URI
url = make_url(connection_string)

vector_store = PGVectorStore.from_params(
    database=url.database,
    host=url.host,
    password=url.password,
    port=url.port, # type: ignore
    user=url.username,
    table_name="embeddings",
    embed_dim=512,  # 'BAAI/bge-small-zh-v1.5' dimension
)


def get_vectorstore():
    return vector_store
