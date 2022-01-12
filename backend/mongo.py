from pymongo import MongoClient
import os

def get_database():

    # Provide the mongodb atlas url to connect python to mongodb using pymongo
    # CONNECTION_STRING = os.environ.get('MONGO_URL')

    CONNECTION_STRING = "mongodb://localhost:27017"

    print("connection url: ", CONNECTION_STRING)

    try:
        # Create a connection using MongoClient. You can import MongoClient or use pymongo.MongoClient
        client = MongoClient(CONNECTION_STRING, serverSelectionTimeoutMS=1000, connectTimeoutMS=20000)
        client.server_info()
        return client['cv']
    except Exception as e:
        print(f"Could not connect to mongodb: {e}")
        exit(1)
