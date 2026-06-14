from bs4 import BeautifulSoup
from dotenv import load_dotenv
from pymongo import MongoClient
import requests
import json
import os

def scrapePost():
    pageURL = "https://www.warhammer-community.com/en-gb/articles/hjdxp86f/saturday-pre-orders-warhammer-40000-armageddon/"
    page = requests.get(pageURL)

    soup = BeautifulSoup(page.text, "html.parser")

    products = soup.find_all("span", attrs = {"class":"font-bold"})

    fmtProducts = []
    for product in products:
        fmtProducts.append(product.text)

    date = soup.find("time", attrs = {"class":"copy-bitter-xs whitespace-nowrap hidden @[175px]/video:block @[225px]/article:block xl:!block"}).text.strip()


    MONGO_URI = os.getenv("MONGO_URI")
    client = MongoClient(MONGO_URI)
    db = client["posts"]
    collection = db["posts"]

    data = {
        "products" : fmtProducts,
        "date" : date,
        "source" : pageURL,
    }

    collection.insert_one(data)


    
load_dotenv()
scrapePost()

    