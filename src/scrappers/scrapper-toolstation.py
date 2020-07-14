import bs4 
from urllib.request import urlopen as uReq
from bs4 import BeautifulSoup as soup

my_url="https://www.toolstation.com/electrical/cables-flexes/c1118?page=2"

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
  results = page_soup.findAll("div",{"class":"ais-hits--item flex-3"})

  for drs in results:
    Product_name= drs.findAll("span",{"class":"f4 f-medium"})[0].contents[0].strip()
    Product_code=  drs.findAll("span",{"class":"sp-product-code f-medium"})[0].contents[1].strip()
    Product_price=drs.findAll("span",{"class":"sp-price f-medium"})[0].contents[0]

    print("Product name: {0}\nProduct code: {1}\nProduct Price: {2}\n***\n\n".format(Product_name, Product_code,Product_price))


if __name__ == "__main__":
  getProductInfo(my_url)
