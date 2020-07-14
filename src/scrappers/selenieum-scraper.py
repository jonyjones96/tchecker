from selenium import webdriver

driver = webdriver.Firefox()
url = "https://www.screwfix.com/c/electrical-lighting/cable/cat8960001#category=cat8960001&page_size=100&page_start=0"
driver.get(url)

# Search for the products
elements = driver.find_elements_by_xpath('//*[contains(@id,"product_box")]')

# Create an array to contain the json requests
products = []

for i in elements:
  items= i.text.split('\n')
  productId=items[0]
  price=items[8]
  products.append({"name":productId,"price": price[1:]})
  #print("Product: {0},\nPrice: {1}\n***\n".format(productId,price))

# Close the connection
driver.close()

# Print out products
print(products)
