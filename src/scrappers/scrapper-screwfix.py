import bs4 
from urllib.request import urlopen as uReq
from bs4 import BeautifulSoup as soup

my_url="https://www.screwfix.com/c/electrical-lighting/cable/cat8960001#category=cat8960001&page_size=10&page_start=1"

def getHtml(url):

  # Opening up connection, grabbing the page
  uClient = uReq(url)
  page_html = uClient.read()
  uClient.close()

  # html Parsing
  page_soup = soup(page_html, "html.parser")

  return  page_soup

def getProductInfo(url):
  page_soup = getHtml(url)
  results = page_soup.findAll("div",{"class":"lg-12 md-24 sm-24 cols"})

  for drs in results:
    Product_name= drs.a.contents[0].strip()
    Product_code=  drs.span.contents[0].strip()
    Product_price=drs.h4.contents[0].strip()

    print("Product name: {0}\nProduct code: {1}\nProduct Price: {2}\n***\n\n".format(Product_name, Product_code,Product_price))


if __name__ == "__main__":
  getProductInfo(my_url)
