import bs4
from urllib.request import urlopen as uReq
from bs4 import BeautifulSoup as soup

my_url="http://www.weathersby.uk.com/site/go/search?sales=true&items=27&includeUnavailable=true"

def getHtml(url):

  # Opening up connection, grabbing the page
  uClient = uReq(url)
  page_html = uClient.read()
  uClient.close()

  # html Parsing
  page_soup = soup(page_html, "html.parser")

  return  page_soup

def propertyInfo(url):
  # The URL directs us to proeprties summary page - this will contain the bulk information on the property
  # Need to get the page html and parse the information through bs4

  page_soup = getHtml(url)
  containers = page_soup.findAll("div",{"id":"particulars"})
  propertyInfo = containers[0].h1
  bedRoomType = propertyInfo.findAll("span",{"class":"bedroomsType"})[0].contents[0].strip()
  address = propertyInfo.findAll("span",{"class":"propertyAddress"})[0].contents[0].strip()
  price = propertyInfo.findAll("span",{"class":"price"})[0].contents[0].strip()
  
  print(f"Bedroom type: {bedRoomType}. Address: {address}. Price: {price}")

  containers=page_soup.findAll("div",{"id":"particularsLeftPanel"})
  summary = containers[0].p.contents[0].strip()
  
  # Get the key Features  
  keyFeatures = []
  feature = type(containers[0].ul.li) # This method is a hack to sepearate useless tags such as '\n'
  for i in containers[0].ul:
    if type(i) == feature:
      keyFeatures.append(i)  

  print(f"Summary: {summary}")

  for i in keyFeatures:
    j = i.contents[1].contents[0]
    print(f"- {j}")

# page_soup = soup(page_html, "html.parser")
page_soup = getHtml(my_url)

# grabs each property
containers = page_soup.findAll("div",{"class":"col-sm-4"})

# Delete this line in future!
# containers = containers[0:1]

for container in containers:
  link = container.a.img["alt"]
  isItSold = container.findAll("div",{"class":"searchResultPhoto"})
  print("Property Name: ",link)
  print("URL: ", container.a["href"])
  
  if bool(isItSold[0].div):
    print("Status: Sold")
  else:
    print("Status: Not Sold")
   
  url = "http://www.weathersby.uk.com" + container.a["href"]
  propertyInfo(url)
  print("\n")
'''  
if __name__ == "__main__":
  main()
'''

