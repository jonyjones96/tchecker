from selenium import webdriver
import time
import requests

def scrapper():
    driver = webdriver.Firefox()
    url = "https://www.screwfix.com/c/electrical-lighting/cable/cat8960001#category=cat8960001&page_size=10&page_start=0"
    driver.get(url)
    
    time.sleep(5)
    
    # Search for the products
    elements = driver.find_elements_by_xpath('//*[contains(@id,"product_box")]')
    
    # Create an array to contain the json requests
    products = []
    
    for i in elements: 
        toolJson = dict()
        toolJson["productFeatures"] = []
        for count, j in enumerate(i.text.split('\n')):
            if count == 0:
                toolJson["productName"] = j
                # The product code is given at the end of the string in between parenthesis.
                toolJson["productCode"] = j[j.find("(")+1:j.find(")")]
            elif j[0] == "£":
                toolJson["productPrice"] = float(j[1:len(j)])
                break
            elif count > 4 and (j.strip() != "INC VAT" or j.strip() != "Click & Collect" or j.strip() != "Deliver"):              
                toolJson["productFeatures"].append(j)
     #   print("Tool = ", toolJson, "\n")
        products.append(toolJson)
        
    # Close the connection
    driver.close()
    
    # Print out products
    print(len(products))

    return products

def sendJson(dataJson):
    print("Sending data...")
    api = "http://localhost:8080/api"
    apiheaders = {'Content-Type':'application/json'}
    for i in dataJson:
        print(i)
        requestData = requests.post(url=api, headers=apiheaders, json=i)



def main():
    # Scrape the data
    dataJson = scrapper()
    # Send the data to api
    sendJson(dataJson)

if __name__ == "__main__":
    main()
