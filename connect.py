from pymongo import MongoClient
import time
import logging

logging.debug("STarting Connection")

print("Waiting...")
#time.sleep(30)

print("Starting connection...")
# Connection details
client = MongoClient('scraper',27017)

db = client.test_database

collection = db.test_collection

# Insert data
post = {"author":"mike"}

posts = db.posts
print("Insterting data")
post_id = posts.insert_one(post).inserted_id

print("Successfully connected to Mongo DB")
