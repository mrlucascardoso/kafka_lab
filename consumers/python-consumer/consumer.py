from kafka import KafkaConsumer
from time import sleep
import os

# Wait for kafka to be ready
sleep(20)

# Create an instance of the Kafka consumer
consumer = KafkaConsumer(os.environ['KAFKA_TOPIC'], bootstrap_servers='kafka:29092', group_id='pythonConsumerGroup')

# Simple print the value. Mind the msg is a tuple with more information.
for msg in consumer:
    print ('Python consumer received value using Hight level API: ' + msg.value)