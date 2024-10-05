import requests
from bs4 import BeautifulSoup

def simple_crawler(url):
    response = requests.get(url)

    if response.status_code == 200:
        # Parse the page content
        soup = BeautifulSoup(response.text, 'html.parser')
        title = soup.title.text  # Extract the title from the page
        print(f'Title: {title}')
    else:
        print(f'Error: Failed to fetch {url}')

# Check the variable and call the function properly
namee = "mainee"  # Define the variable

if namee == "mainee":
    sample_url = 'https://google.com'
    simple_crawler(sample_url)
