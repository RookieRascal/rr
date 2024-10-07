use std::collections::HashMap;
use serde::{Serialize, Deserialize};

#[derive(Debug, Serialize, Deserialize)]
struct Document {
    id: usize,
    content: String,
}

struct Indexer {
    index: HashMap<String, Vec<usize>>,
}

impl Indexer {
    fn new() -> Self {
        Indexer {
            index: HashMap::new(),
        }
    }

    async fn add_document(&mut self, doc: Document) {
        let words = doc.content.split_whitespace();
        for word in words {
            self.index.entry(word.to_string())
                .or_insert(vec![])
                .push(doc.id);
        }
    }

    fn search(&self, query: &str) -> Vec<usize> {
        self.index.get(query).cloned().unwrap_or(vec![])
    }
}

#[tokio::main]
async fn main() {
    let mut indexer = Indexer::new();
    let doc = Document { id: 1, content: "rust is fast and safe".to_string() };
    indexer.add_document(doc).await;
    let results = indexer.search("rust");
    println!("Search results: {:?}", results);
}










































