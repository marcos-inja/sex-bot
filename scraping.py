import requests
import sys
import concurrent.futures
from threading import Thread
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


    
def get_page_joke(init: str) -> list:
    joke_list = get_links(URL_BASE + init)
    with concurrent.futures.ThreadPoolExecutor() as executor:
        file = []
        jokes = []
        for joke_link in joke_list:
            file.append(
                executor.submit(
                    get_joke, URL_BASE + joke_link
                )
            )
        for future in concurrent.futures.as_completed(file):
            jokes.append(future.result())

    return jokes


def get_links_all_categories() -> list:
    page = get_url_soup(URL_BASE)

    links = page.find(class_ = "menuEsq")
    tag_a = links.find_all("a")
    link_categories = []
    for link in tag_a:
        link_categories.append(link.attrs["href"])
    return link_categories


links_categories = get_links_all_categories()


with concurrent.futures.ThreadPoolExecutor() as executor:
    file = []
    for link in links_categories:
        file.append(
            executor.submit(
                get_page_joke, link
            )
        )

    for future in concurrent.futures.as_completed(file):
        with open(f"jokes.json", "r") as fp:
            jokes_file = json.load(fp)

        for joke in future.result():
            jokes_file.append(joke)

        with open(f"jokes.json", "w") as fp:
            json.dump(jokes_file, fp, indent=4, ensure_ascii=False)