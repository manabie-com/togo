FROM alpine:3.14.5
LABEL Maintainer="Han" \
      Description="docker"

# Install packages
RUN apk --no-cache add php7 php7-fpm php7-mysqli php7-json php7-openssl php7-curl php7-pdo_mysql php7-xmlwriter php7-xsl php7-zip\
    php7-zlib php7-xml php7-phar php7-intl php7-dom php7-xmlreader php7-ctype php7-session php7-redis php7-simplexml php7-soap \
    php7-mbstring php7-fileinfo php7-gd php7-pear php7-dev php7-pecl-xdebug php7-bcmath php7-tokenizer nginx supervisor curl composer

ENV WORKDIR=/var/www/html

WORKDIR  $WORKDIR

# Set permission
RUN chown -R nobody.nobody /run && \
  chown -R nobody:nobody /var/lib/nginx && \
  chown -R nobody:nobody /var/log/nginx

RUN chown -R nobody:nobody $WORKDIR

# Make the document root a volume
VOLUME $WORKDIR

USER nobody
# Expose the port nginx is reachable on
EXPOSE 8080

# Let supervisord start nginx & php-fpm
CMD ["/usr/bin/supervisord", "-c", "/etc/supervisor/conf.d/supervisord.conf"]

# Configure a healthcheck to validate that everything is up&running
HEALTHCHECK --timeout=10s CMD curl --silent --fail http://127.0.0.1:8080/fpm-ping
