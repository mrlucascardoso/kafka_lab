from kafka import KafkaProducer
from time import sleep
import os
import random

# Wait for kafka to be ready, but get ready before others.
sleep(15)

# Create an instance of the Kafka producer
producer = KafkaProducer(bootstrap_servers='kafka:29092',
                            value_serializer=lambda v: str(v).encode('utf-8'))

# Let's create some random data to publish
while True:
    var = random.randint(1,999)
    print("Publishing value ", var, " on ", os.environ['KAFKA_TOPIC'])
    producer.send(os.environ['KAFKA_TOPIC'], var)
    sleep(1)