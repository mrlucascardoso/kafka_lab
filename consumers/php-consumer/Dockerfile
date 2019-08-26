FROM php:7.2-cli

ADD . .

RUN apt-get update && apt install -y librdkafka-dev && pecl install rdkafka

RUN mv "$PHP_INI_DIR/php.ini-development" "$PHP_INI_DIR/php.ini"

RUN echo "extension=rdkafka.so" > $PHP_INI_DIR/conf.d/rdkafka.ini

CMD php consumer.php
