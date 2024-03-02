import os
from dotenv import load_dotenv
from ultralytics import YOLO
import cv2

load_dotenv()

modelPath = os.getenv("ML_MODEL")
if modelPath == None:
    print("No trained model found to import")
    os._exit(1)

# Load a model
model = YOLO(modelPath)

def myPred(path):
    im2 = cv2.imread(path)
    results = model(source=im2)
    return results
