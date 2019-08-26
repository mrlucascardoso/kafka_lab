<?php

// Wait for kafka to be ready
sleep(20);

if (getenv("KAFKA_API_LEVEL") == 'High') {
    $conf = new RdKafka\Conf();
    $conf->set('group.id', 'phpConsumerGroup'); // Group ID must be set, consumers with same group.id consume different partitions.
    $conf->set('metadata.broker.list', "kafka:29092"); // Broker address.

    $consumer = new RdKafka\KafkaConsumer($conf); // This is a High Level API.
    $consumer->subscribe([getenv("KAFKA_TOPIC")]);

    while (true) {
        $message = $consumer->consume(3000); // The parameter we send is the max timeout in ms.
        var_dump("PHP consumer received value using High level API: " . $message->payload); // Just print the message payload, message have some more information about the received data.
    }
}

if (getenv("KAFKA_API_LEVEL") == 'Low') {
    $conf = new RdKafka\Conf(); // Since this is a minimum configuration state, there's not much to add here, as in a Low Level API the consumer has to know everything about the topic.

    $consumer = new RdKafka\Consumer($conf); // This is a Low Level API
    $consumer->addBrokers("kafka:29092"); // See that this time the consumer has to know the kafka address, this is no longer a global configuration.

    $topic = $consumer->newTopic(getenv("KAFKA_TOPIC"));
    $topic->consumeStart(0, 0); // First parameter is the partition index and second is the offset.

    while (true) {
        // The first argument is the partition (again).
        // The second argument is the timeout in ms.
        $message = $topic->consume(0, 3000);
        var_dump("PHP consumer received value using Low level API: " . $message->payload); // Just print the message payload, message have some more information about the received data.
    }
}