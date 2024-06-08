from fastapi import FastAPI
from typing import Optional, List
from llama_index.core import VectorStoreIndex
from llama_index.core.llms import ChatMessage, MessageRole
from llama_index.core.vector_stores.types import VectorStoreQuery
from pydantic import BaseModel, Field

from noteAi.rag import RAG
from noteAi.pg_vecstore import get_vectorstore


app = FastAPI()


@app.get("/")
def root():
    return {"message": "Research RAG"}


class ResponseQuestion(BaseModel):
    search_result: str
    sources: List[str]


class SourceModal(BaseModel):
    id: str
    title: Optional[str]
    text: str


class ResponseSearch(BaseModel):
    sources: List[SourceModal]


class QuerySearch(BaseModel):
    query: str
    similarity_top_k: Optional[int] = Field(default=1, ge=1, le=5)


class ChatSourceModal(BaseModel):
    id: str
    title: Optional[str]
    score: Optional[float]


class ResponseChat(BaseModel):
    content: str
    sources: List[ChatSourceModal]


class ChatMessageModel(BaseModel):
    role: MessageRole
    content: str


class QueryChat(BaseModel):
    query: str
    similarity_top_k: Optional[int] = Field(default=1, ge=1, le=5)
    history: Optional[List[ChatMessageModel]] = None


@app.post("/ai/question", response_model=ResponseQuestion, status_code=200)
def question(user_id: str, query: QuerySearch):
    rag = RAG(userid=user_id)
    vector_store = get_vectorstore()
    index = VectorStoreIndex.from_vector_store(vector_store)
    query_engine = index.as_query_engine(
        similarity_top_k=query.similarity_top_k,
        output=ResponseQuestion,
        verbose=True,
        filters=rag.user_selector,
    )
    response = query_engine.query(query.query)
    response_object = ResponseQuestion(
        search_result=str(response).strip(),
        sources=[node.metadata.get("id") for node in response.source_nodes],  # type: ignore
    )
    return response_object


@app.post("/ai/chat", response_model=ResponseChat, status_code=200)
def chat(user_id: str, query: QueryChat):
    rag = RAG(userid=user_id)
    vector_store = get_vectorstore()
    index = VectorStoreIndex.from_vector_store(vector_store)
    chat_engine = index.as_chat_engine(
        similarity_top_k=query.similarity_top_k,
        verbose=True,
        filters=rag.user_selector,
    )
    response = chat_engine.chat(
        query.query,
        chat_history=[
            ChatMessage.from_str(content=chat.content, role=chat.role)
            for chat in query.history or []
        ],
    )
    return ResponseChat(
        content=str(response).strip(),
        sources=[
            ChatSourceModal(
                id=node.metadata.get("id"),  # type: ignore
                title=node.metadata.get("title"),
                score= node.get_score(),
            )
            for node in response.source_nodes  # type: ignore
        ],
    )


@app.post("/ai/search", response_model=ResponseSearch, status_code=200)
def search(user_id: str, query: QuerySearch):
    rag = RAG(userid=user_id)
    embedding = rag.get_query_embedding(query.query)
    vector_store_query = VectorStoreQuery(
        query_embedding=embedding,
        query_str=query.query,
        similarity_top_k=query.similarity_top_k,  # type: ignore
        filters=rag.user_selector,
    )
    vector_store = get_vectorstore()
    result = vector_store.query(query=vector_store_query)
    response_object = ResponseSearch(
        sources=[
            SourceModal(
                id=node.metadata.get("id"),  # type: ignore
                text=node.get_content(),
                title=node.metadata.get("title"),
            )
            for node in result.nodes  # type: ignore
        ]
    )
    return response_object


@app.post("/ai/refresh", status_code=200)
def refresh(user_id: str):
    rag = RAG(userid=user_id)
    rag.refresh_docs()
    return {"message": "Refreshed successfully!"}
