This is a simple test to evaluate communication between different languages using kafka as a broker.

On Go we are using "kafka-go" (https://github.com/segmentio/kafka-go), which seems to be the best lib to be used, "confluent-kafka-go" also seems to be good (it uses librdkafka which is a C lib). The libs seems nice to use, messages are consumed in chuncks, which seems to lead to less network requests, i believe that with some fine tuning this lib may provide the best resource optimization.

On PHP we are using "php-rdkafka" (https://github.com/arnaud-lb/php-rdkafka), which uses librdkafka, a C lib. Even though the lib seems to offer everything that Kafka has to offer, it might be challeging to config and tune it, the documentation is kinda poor and some configs have to be guessed, noneless, seems to work fine. A strange behavior is that the consumer seems to struggle to reconnect with Kafka API, maybe because it enforces the use of Group ID, still gotta go further on that.

Python lib is a no brainer, copy & paste and things work, the used lib was "kafka-python" (https://github.com/dpkp/kafka-python), it implies a lot of configs and low-level API use seems to be discouraged.

Java lib is to be tested, since it has an official client it has a lower priority.

To test the scenario, simply run:
docker-compose up or docker-compose up -d
-d stands for detached, if you use this you will need to get containers logs with docker logs <container-id>.