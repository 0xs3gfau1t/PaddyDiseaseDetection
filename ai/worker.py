import os
from http import client
from typing import Callable

from db import DbClient
from paddyinference import myPred

class Worker:
    """
    Total 3 database connections
    """
    db = DbClient
    toIdentify: str
    host: str
    tempFile: str

    def __init__(self, id: str):
        self.toIdentify = id

        host = os.getenv("SUPABASE_HOST")
        if host == None:
            print("No supabase host env configured")
            os._exit(1)
        self.host = host

    def run(self, done: Callable, ack: Callable, method):
        # Synchronous high time/resource consuming operation
        diseaseName = self.classify()
        if diseaseName != None:
            diseaseId = self.db().getDiseaseFromName(diseaseName)
            if diseaseId != None:
                self.update(diseaseId[0])
                print("Disease: ", diseaseName)
            else:
                print("[x] Couldn't find disease with name ", diseaseName)

            ack(delivery_tag=method.delivery_tag)
        done() # Release the lock

    def classify(self):
        imageIdentifier = self.db().getImageIdentifier(self.toIdentify)
        if imageIdentifier != None: 
            tempFile = self.getImagePath(imageIdentifier)
            if tempFile != None:
                self.tempFile = tempFile
                return myPred(self.tempFile)

    def getImagePath(self, identifier):
        conn = client.HTTPSConnection(self.host)
        conn.request("GET",
                     self.constructImageLink(identifier), headers={"Host": self.host})
        response = conn.getresponse()
        if response.getcode() != 200:
            return None

        fName = self.toIdentify+"-image"
        with open(fName, "wb") as file:
            file.write(response.read())
        return fName

    def constructImageLink(self, identifier):
        return "/storage/v1/object/public/{bucket}/{identifier}".format(
                         bucket=os.getenv("IMAGE_BUCKET"),
                         identifier=identifier[0])

    def update(self, diseaseId):
        self.db().updateDone(self.toIdentify, diseaseId)

    def __del__(self):
        if self.tempFile != None:
            os.remove(self.tempFile)

if __name__ == "__main__":
    Worker("8210287b-1e2a-418d-8f89-433f99d72a2d").classify()
