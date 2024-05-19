from typing import Optional
import os
from llama_index.core import VectorStoreIndex
from llama_index.embeddings.fastembed import FastEmbedEmbedding
from llama_index.llms.together import TogetherLLM
from llama_index.core import load_index_from_storage, StorageContext, Settings
from llama_index.storage.docstore.postgres import PostgresDocumentStore
from llama_index.core.vector_stores.types import ExactMatchFilter, MetadataFilters


from noteAi.config import AppConfig
from noteAi.data_loader import CouchDBReader
from noteAi.pg_vecstore import get_vectorstore

Settings.embed_model = FastEmbedEmbedding(
    model_name="BAAI/bge-small-zh-v1.5", cache_dir="./model_cache"
)

Settings.llm = TogetherLLM(
    model="Qwen/Qwen1.5-14B-Chat", api_key=AppConfig.TOGETHER_API_KEY
)

reader = CouchDBReader(
    user=AppConfig.COUCHDB_USER,
    pwd=AppConfig.COUCHDB_PASSWORD,
    host=AppConfig.COUCHDB_HOST,
    port=AppConfig.COUCHDB_PORT,
)

COUCHDB_QUERY = {"_id": {"$gt": None}}


class RAG:
    def __init__(self, userid: str, dbname: Optional[str] = None) -> None:
        self.userid = userid
        self.dbname = f"userdb-{userid}"
        if dbname:
            self.dbname = dbname

    @property
    def user_selector(self) -> MetadataFilters:
        return MetadataFilters(
            filters=[ExactMatchFilter(key="user", value=self.dbname)]
        )

    @staticmethod
    def get_query_embedding(query: str):
        return Settings.embed_model.get_query_embedding(query)

    def refresh_docs(self) -> None:
        vector_store = get_vectorstore()

        docstore = PostgresDocumentStore.from_uri(
            uri=AppConfig.POSTGRES_URI,
            namespace=f"userdb-{self.userid}",
        )
        documents = reader.load_data(self.dbname, query=COUCHDB_QUERY)

        # Incremental computation of the index
        if os.path.exists("./persist"):
            storage_context = StorageContext.from_defaults(
                vector_store=vector_store,
                docstore=docstore,
                persist_dir="./persist",
            )
            index = load_index_from_storage(storage_context=storage_context)
            index.refresh_ref_docs(documents=documents)

        else:
            storage_context = StorageContext.from_defaults(
                vector_store=vector_store, docstore=docstore
            )
            index = VectorStoreIndex.from_documents(
                documents,
                show_progress=True,
                storage_context=storage_context,
            )

        # Remove documents that are no longer user's notes
        exist_id_set = set()
        for doc in documents:
            exist_id_set.add(doc.doc_id)

        docstore_id_set = index.docstore.get_all_document_hashes().values()
        for id in docstore_id_set:
            if id not in exist_id_set:
                index.delete_ref_doc(ref_doc_id=id, delete_from_docstore=True)

        index.storage_context.persist(persist_dir="./persist")
