import json
import os
import requests

from .paddyinference import myPred

class Worker:
    tempFile = None

    def __init__(self, parsedData):
        self.toIdentify = parsedData.get("id")
        self.link = parsedData.get("link")

    def run(self) -> str:
        # Synchronous high time/resource consuming operation
        response = requests.get(self.link)
        disease = "unknown"
        if response.status_code == 200:
            fName = self.toIdentify+"-image"
            with open(fName, "wb") as file:
                if response._content:
                    file.write(response._content)

            self.tempFile = fName
            disease = myPred(self.tempFile)
        return json.dumps({"id": self.toIdentify, "disease": disease, "status": "processed"})

    def __del__(self):
        if self.tempFile != None:
            os.remove(self.tempFile)
