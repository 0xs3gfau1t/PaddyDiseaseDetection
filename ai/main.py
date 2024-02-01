import json
import time
import os
import threading
import pika
from pika.adapters.select_connection import SelectConnection
from pika.channel import Channel
from pika.connection import Connection
from dotenv import load_dotenv
from worker import Worker

load_dotenv()

class Supervisor:
    

    conn: SelectConnection 
    channelConsumer: Channel
    rabbitQueueConsumer: str
    rabbitQueueProducer: str
    
    def __init__(self, workers=5):
        self.nWorkers = workers

        rabbitHost = os.getenv("RABBIT_HOST")
        rabbitPort = os.getenv("RABBIT_PORT")
        rabbitUser = os.getenv("RABBIT_USER")
        rabbitPass = os.getenv("RABBIT_PASS")
        rabbitQueueC = os.getenv("RABBIT_QUEUE_CONSUMER")
        rabbitQueueP = os.getenv("RABBIT_QUEUE_PRODUCER")
        maxNWorkers = os.getenv("MAX_ML_WORKERS")
        if not rabbitPort or not rabbitUser or not rabbitPass or not rabbitHost or not rabbitQueueC or not rabbitQueueP or not maxNWorkers:
            print("[x] No rabbit url found in env. Exiting")
            os._exit(1)

        self.maxNWorkers = int(maxNWorkers)
        self.workerThreadsLock = threading.Semaphore(int(maxNWorkers))


        self.rabbitQueueConsumer = rabbitQueueC
        self.rabbitQueueProducer = rabbitQueueP

        # Step #1
        pikaConn = pika.URLParameters(f"amqp://{rabbitUser}:{rabbitPass}@{rabbitHost}:{rabbitPort}")

        self.conn = pika.SelectConnection(parameters=pikaConn,
                                          on_open_callback=self.onConnected,
                                          on_open_error_callback=self.onConnectedError,
                                          on_close_callback=self.onClose,
                                          )

        try:
            self.mainLoop = threading.Thread(target=self.conn.ioloop.start, daemon=True)
        except KeyboardInterrupt:
            print("Closing connection")
            # Gracefully close the connection
            self.conn.close()
            # Loop until we're fully closed.
            # The on_close callback is required to stop the io loop
            self.conn.ioloop.start()

    def run(self):
        self.mainLoop.start()
        self.monitor()

    # Step #2
    def onConnected(self, connection):
        connection.channel(on_open_callback=self.onChannelOpen)

    def onConnectedError(self, conn:Connection, excep):
        print("Error connecting: ", conn)

    # Step #3
    def onChannelOpen(self, new_channel:Channel):
        """Called when our channel has opened"""
        self.channelConsumer = new_channel
        self.channelConsumer.queue_declare(queue=self.rabbitQueueConsumer, durable=True, exclusive=False, auto_delete=False, callback=self.onQueueDeclaredConsumer)

    def onClose(self, connection, exception):
        connection.ioloop.close()

    # Step #4
    def onQueueDeclaredConsumer(self, frame):
        """Called when RabbitMQ has told us our Queue has been declared, frame is the response from RabbitMQ"""
        self.channelConsumer.basic_consume(self.rabbitQueueConsumer, self.handleDelivery)

    # Step #5
    def handleDelivery(self, channel:Channel, method, header, body):
        """Called when we receive a message from RabbitMQ"""
        try:
            parsedData = json.loads(bytes.decode(body, encoding='utf-8'))
            responseMessage = {"id": parsedData.get("id"), "disease": "unknown", "status": "processing"}
            # self.respond(json.dumps(responseMessage))
        except:
            print("Couldn't parse data")
            return

        self.workerThreadsLock.acquire()
        try:
            self.respond(Worker(parsedData).run())
        except Exception as e:
            print("Error: ", e)
            responseMessage["status"] = "failed"
            self.respond(json.dumps(responseMessage))

        self.workerThreadsLock.release()
        channel.basic_ack(delivery_tag=method.delivery_tag)

    def respond(self, msg:str):
        self.channelConsumer.basic_publish(
                body=msg,
                exchange="",
                routing_key=self.rabbitQueueProducer
                )

    def monitor(self):
        while True:
            activeWorkers = self.maxNWorkers - self.workerThreadsLock._value
            print(f"{activeWorkers} workers on the job", end="\r", flush=True)
            if not self.mainLoop.is_alive():
                print("Main worker dead. Reviving...")
                self.mainLoop.run()
            time.sleep(1) 

    def __del__(self):
        if self.conn != None:
            self.conn.close()

Supervisor().run()
