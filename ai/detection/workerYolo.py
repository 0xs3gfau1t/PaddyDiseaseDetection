import json
import os
import requests
from .paddyinferenceYolo import myPred
import random

colors = [[random.randint(0, 255) for _ in range(3)] for _ in range(24)]

class Worker:
    tempFile = None

    def __init__(self, parsedData):
        self.toIdentify = parsedData.get("id")
        self.link = parsedData.get("link")

    def run(self) -> str:
        # Synchronous high time/resource consuming operation
        response = requests.get(self.link)
        r = []
        if response.status_code == 200:
            fName = self.toIdentify+"-image"
            with open(fName, "wb") as file:
                if response._content:
                    file.write(response._content)

            self.tempFile = fName
            result = myPred(self.tempFile)[0]

            data = result.boxes.data.cpu().tolist()
            h,w = result.orig_shape
            names = result.names

            for row in data:
                box = [row[0] / w, row[1] / h, row[2] / w, row[3] / h]
                conf = row[4]
                classId = int(row[5])
                name = names[classId]
                r.append(
                    {
                        "box": box,
                        "confidence": conf,
                        "classId": classId,
                        "name": name,
                        "color": "#%02x%02x%02x" % tuple(colors[classId]),
                    }
                )
        return json.dumps({"frames": r, "id": self.toIdentify})

    def __del__(self):
        if self.tempFile != None:
            os.remove(self.tempFile)
