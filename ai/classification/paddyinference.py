import os
from dotenv import load_dotenv
import torch
from torchvision.io import read_image
from torchvision import transforms

load_dotenv()

diseases = ['bacterial_leaf_blight',
            'bacterial_leaf_streak',
            'bacterial_panicle_blight',
            'black_stem_borer',
            'blast',
            'brown_spot',
            'downy_mildew',
            'hispa',
            'leaf_roller',
            'normal',
            'tungro',
            'white_stem_borer',
            'yellow_stem_borer'
          ]

model = torch.hub.load('pytorch/vision:v0.10.0', 'resnet18', weights=None)
model.eval()


modelPath = os.getenv("ML_MODEL")
if modelPath == None:
    print("No trained model found to import")
    os._exit(1)


model.load_state_dict(torch.load(modelPath, map_location=torch.device('cpu')))

transform2 = transforms.Compose([
    transforms.ToPILImage(),
    transforms.Resize((256,256)),
    transforms.ToTensor(),
])

def myPred(path):
    im = read_image(path)[:3]
    im = transform2(im).unsqueeze(0)
    pred = model(im)
    dis = diseases[pred.argmax().tolist()]
    return dis
