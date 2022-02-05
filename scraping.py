import requests
import sys
import concurrent.futures
import json

from bs4 import BeautifulSoup


URL_BASE = "https://www.piadasnet.com/"

def get_url_soup(url: str) -> BeautifulSoup:
    try:
        page_source = requests.get(url)
    except Exception as err:
        sys.exit(str(err))
    
    soup = BeautifulSoup(page_source.text, "html.parser")
    return soup


def get_links(url: str) -> list:
    page = get_url_soup(url)
    joke_div = page.find_all(class_ = "linkvisitado")
    joke = []
    for link in joke_div:
        joke.append(link.find("a").attrs["href"])
        
    return joke


def get_joke(url: str) -> str:
    page = get_url_soup(url)
    joke = page.find(class_ = "piada").text
    joke = joke.replace("\r\n", " ").replace("\n", " ").replace("\r", " ")
    return joke

with open("jokes.json", "r") as fp:
    jokes = json.load(fp)
    

def get_page_joke(init: str) -> None:
    joke_list = get_links(URL_BASE + init)
    with concurrent.futures.ThreadPoolExecutor() as executor:
            file = []
            for joke_link in joke_list:
                file.append(
                    executor.submit(
                        get_joke, URL_BASE + joke_link
                    )
                )
            for future in concurrent.futures.as_completed(file):
                jokes.append(future.result())


PAGE_INIT = "piadas-de-casais.htm"
get_page_joke(PAGE_INIT)
    
with open(f"jokes.json", "w") as fp:
    json.dump(jokes, fp, indent=4, ensure_ascii=False)