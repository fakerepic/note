"""CouchDB client."""

# import json
import logging
from typing import Dict, List, Optional

import couchdb3
from llama_index.core.readers.base import BaseReader
from llama_index.core.schema import Document


class CouchDBReader(BaseReader):
    """Simple CouchDB reader.

    Concatenates each CouchDB doc into Document used by LlamaIndex.

    Args:
        couchdb_url (str): CouchDB Full URL.
        max_docs (int): Maximum number of documents to load.

    """

    def __init__(
        self,
        user: str,
        pwd: str,
        host: str,
        port: int,
        couchdb_url: Optional[str] = None,
        max_docs: int = 1000,
    ) -> None:
        """Initialize with parameters."""
        if couchdb_url is not None:
            self.client = couchdb3.Server(couchdb_url)
        else:
            self.client = couchdb3.Server(f"http://{user}:{pwd}@{host}:{port}")
        self.max_docs = max_docs

    def load_data(self, db_name: str, query: Optional[Dict] = None) -> List[Document]:
        """Load data from the input directory.

        Args:
            db_name (str): name of the database.
            query (Optional[str]): query to filter documents.
                Defaults to None

        Returns:
            List[Document]: A list of documents.

        """
        documents = []
        db = self.client.get(db_name)
        if query is None:
            # if no query is specified, return all docs in database
            logging.debug("showing all docs")
            results = db.view("_all_docs", include_docs=True)
        else:
            logging.debug("executing query")
            results = db.find(query)

        docs = results.get("docs")
        if docs is not None:
            for item in docs:
                # check that the _id field exists
                if "_id" not in item:
                    raise ValueError("`_id` field not found in CouchDB document.")
                if "content" not in item:
                    continue
                documents.append(
                    Document(
                        text=item.get("content"),
                        id_=item.get("_id"),
                        metadata={
                            "id": item.get("_id"),
                            "title": item.get("title"),
                            "user": db_name,
                        },
                    )
                )

        return documents
