import json
import os
import pika
from dotenv import load_dotenv
import random

load_dotenv()
rabbitHost = os.getenv("RABBIT_HOST")
rabbitPort = os.getenv("RABBIT_PORT")
rabbitUser = os.getenv("RABBIT_USER")
rabbitPass = os.getenv("RABBIT_PASS")
rabbitQueueC = os.getenv("RABBIT_QUEUE_CONSUMER")
rabbitQueueP = os.getenv("RABBIT_QUEUE_PRODUCER")
images = [
        {"id": "void", "link": "https://d3i71xaburhd42.cloudfront.net/8761632fb489e99863091d656be5324c0227614b/1-Figure1-1.png"},
        {"id": "void", "link": "https://d3i71xaburhd42.cloudfront.net/8761632fb489e99863091d656be5324c0227614b/1-Figure1-1.png"},
        {"id": "void", "link": "https://d3i71xaburhd42.cloudfront.net/8761632fb489e99863091d656be5324c0227614b/1-Figure1-1.png"},
        ]

if not rabbitPort or not rabbitUser or not rabbitPass or not rabbitHost or not rabbitQueueC or not rabbitQueueP:
    print("[x] No rabbit url found in env. Exiting")
    os._exit(1)

pikaConn = pika.URLParameters(f"amqp://{rabbitUser}:{rabbitPass}@{rabbitHost}:{rabbitPort}")


connection = pika.BlockingConnection(pikaConn)
channel = connection.channel()

channel.queue_declare(queue=rabbitQueueC, durable=True)

for i in range(100):
    image = images[random.randrange(0,2)]
    image['id'] += str(i)
    channel.basic_publish(exchange='',
                          routing_key=rabbitQueueC,
                          body=json.dumps(image)
                          )
    print(" [x] Sent Image")
