This is a simple test to evaluate communication between different languages using kafka as a broker.

What i learned so far:
- We can use High level API or Low level API, in High level API kafka leverage Zookeeper management to keep track of the offsets that each consumer has within each topic its consuming (you must provide a Consumer Group name, which will be used to register which group the application belongs to). In Low level API you must keep track of this on application side, you also have to handle what partition you wish to connect to, which may be tricky on an always changing environment.

- Either you scale or keep message order, you can't have both, since Kafka use partitions to scale and they don't keep track of sent messages, 2 partitions may send messages about the same topic, and one might be faster than the other, sometimes making newer messages reach the application before old messages.

- Kafka uses a lot of features that modern Hard Disks offers, this makes Kafka very resilient and incredibly fast, but it makes the solution very disk dependent which is a challenge on a Cloud environment, Stateful applications are harder to maintain than ephemeral applications or even RAM dependent applications.

- At first i thought that Kafka could be used to supply event sourcing needs, but that's just impraticable, Kafka has no filtering, to replay an entity you would have to receive back ALL messages from a specific topic, also, there's no snapshot feature, this should be implemented by the application, thus forcing 2 different sources for the events.

About the libs:

On Go we are using "kafka-go" (https://github.com/segmentio/kafka-go), which seems to be the best lib to be used, "confluent-kafka-go" also seems to be good (it uses librdkafka which is a C lib). The lib easy to configure, messages are consumed in chuncks, which seems to lead to less network requests, i believe that with some fine tuning this lib may provide the best resource optimization.

On PHP we are using "php-rdkafka" (https://github.com/arnaud-lb/php-rdkafka), which uses librdkafka, a C lib. Even though the lib seems to offer everything that Kafka has to offer, it might be challeging to config and tune it, the documentation is kinda poor and some configs have to be guessed, noneless, seems to work fine. A strange behavior is that the consumer seems to struggle to reconnect with Kafka API, maybe because it enforces the use of Group ID, still gotta go further on that.

Python lib is a no brainer, copy & paste and things work, the used lib was "kafka-python" (https://github.com/dpkp/kafka-python), it implies a lot of configs and low-level API use seems to be discouraged.

Java lib is to be tested, since it has an official client it has a lower priority.

To test the scenario, simply run:
docker-compose up or docker-compose up -d
-d stands for detached, if you use this you will need to get containers logs with docker logs <container-id>.