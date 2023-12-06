import time
import os
import threading
import pika
from pika.adapters.select_connection import SelectConnection
from pika.channel import Channel

from worker import Worker

maxNWorkers = 1

class Supervisor:
    workerThreadsLock = threading.Semaphore(maxNWorkers)

    conn: SelectConnection 
    channel: Channel
    rabbitQueue: str
    
    def __init__(self, workers=5):
        self.nWorkers = workers

        rabbitHost = os.getenv("RABBIT_HOST")
        rabbitPort = os.getenv("RABBIT_PORT")
        rabbitUser = os.getenv("RABBIT_USER")
        rabbitPass = os.getenv("RABBIT_PASS")
        rabbitQueue = os.getenv("RABBIT_QUEUE")
        if not rabbitPort or not rabbitUser or not rabbitPass or not rabbitHost or not rabbitQueue:
            print("[x] No rabbit url found in env. Exiting")
            os._exit(1)

        self.rabbitQueue = rabbitQueue

        # Step #1
        pikaConn = pika.URLParameters(f"amqp://{rabbitUser}:{rabbitPass}@{rabbitHost}:{rabbitPort}")

        self.conn = pika.SelectConnection(parameters=pikaConn,
                                          on_open_callback=self.onConnected,
                                          on_open_error_callback=self.onConnectedError,
                                          on_close_callback=self.onClose,
                                          )

    def run(self):
        mainLoop = threading.Thread(target=self.conn.ioloop.start, daemon=True)
        try:
            self.monitor()
        except KeyboardInterrupt:
            print("Closing connection")
            # Gracefully close the connection
            self.conn.close()
            # Loop until we're fully closed.
            # The on_close callback is required to stop the io loop
            self.conn.ioloop.start()
            mainLoop.join()

    # Step #2
    def onConnected(self, connection):
        connection.channel(on_open_callback=self.onChannelOpen)

    def onConnectedError(self, conn):
        print("Error connecting: ", conn)

    # Step #3
    def onChannelOpen(self, new_channel):
        """Called when our channel has opened"""
        self.channel = new_channel
        self.channel.queue_declare(queue=self.rabbitQueue, durable=True, exclusive=False, auto_delete=False, callback=self.onQueueDeclared)

    def onClose(self, connection, exception):
        connection.ioloop.close()

    # Step #4
    def onQueueDeclared(self, frame):
        """Called when RabbitMQ has told us our Queue has been declared, frame is the response from RabbitMQ"""
        self.channel.basic_consume(self.rabbitQueue, self.handleDelivery)

    # Step #5
    def handleDelivery(self, channel, method, header, body):
        """Called when we receive a message from RabbitMQ"""
        self.workerThreadsLock.acquire()
        decodedBody = bytes.decode(body, encoding='utf-8')
        newWorker = threading.Thread(target=Worker(decodedBody).run, args=(
            self.workerThreadsLock.release,
            channel.basic_ack,
            method
            ))
        newWorker.start()

    def monitor(self):
        while True:
            activeWorkers = maxNWorkers - self.workerThreadsLock._value
            print(f"{activeWorkers} workers on the job")
            time.sleep(1) 

    def __del__(self):
        if self.conn != None:
            self.conn.close()

Supervisor().run()
