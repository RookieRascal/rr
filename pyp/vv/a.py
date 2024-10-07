from collections import defaultdict
import re

class SearchEngineIndexer:
    def __init__(self):
        self.index = defaultdict(list)  # Word -> list of doc_ids
    
    def add_document(self, doc_id, text):
        words = re.findall(r'\w+', text.lower())  # Tokenize text
        for word in set(words):  # Use set to avoid duplicate words in the same doc
            self.index[word].append(doc_id)
    
    def build_index(self, docs):
        for doc_id, content in docs.items():
            self.add_document(doc_id, content)

# Example usage
documents = {
    1: "Search engines are programs that search the web.",
    2: "A web crawler is used to gather data for search engines.",
    3: "Indexing is crucial for search engine efficiency."
}

indexer = SearchEngineIndexer()
indexer.build_index(documents)
print(dict(indexer.index))  # Outputs inverted index
