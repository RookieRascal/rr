use reqwest;
use scraper::{Html, Selector};

async fn simple_crawler(url: &str) {
    // Make a GET request to the given URL
    match reqwest::get(url).await {
        Ok(response) => {
            if response.status().is_success() {
                // Extract the text from the response
                let body = response.text().await.unwrap();
                // Parse the HTML
                let document = Html::parse_document(&body);
                // Create a selector for the <title> tag
                let title_selector = Selector::parse("title").unwrap();
                
                // Find and print the title of the page
                if let Some(title_element) = document.select(&title_selector).next() {
                    let title = title_element.text().collect::<Vec<_>>().concat();
                    println!("Title: {}", title);
                } else {
                    println!("Error: No title found on the page");
                }
            } else {
                println!("Error: Failed to fetch {}", url);
            }
        },
        Err(e) => {
            println!("Error: {}", e);
        }
    }
}

#[tokio::main]
async fn main() {
    let namee = "mainee"; // Define a variable
    if namee == "mainee" {
        let sample_url = "https://google.com";
        simple_crawler(sample_url).await;
    }
}
