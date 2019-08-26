<?php

// Wait for kafka to be ready
sleep(20);

// Miminum config for consumer
$conf = new RdKafka\Conf();
$conf->set('log_level', 6);
$conf->set('group.id', 'phpConsumerGroup'); // Group ID must be set, consumers with same group.id consume different partitions.
$conf->set('metadata.broker.list', "kafka:29092"); // Broker address

$consumer = new RdKafka\KafkaConsumer($conf);
$consumer->subscribe([getenv("KAFKA_TOPIC")]);

while (true) {
    $message = $consumer->consume(3000); // The parameter we send is the max timeout in ms.
    var_dump($message->payload); // Just print the message payload, message have some more information about the received data.
}