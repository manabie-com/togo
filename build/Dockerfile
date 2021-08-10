FROM ubuntu:18.04

# Fixes some weird terminal issues such as broken clear / CTRL+L
ENV TERM=linux \
    php_ini_file="/etc/php/7.3/fpm/php.ini"    

RUN export DEBIAN_FRONTEND=noninteractive \
    && apt-get update \
    && apt-get install -y --no-install-recommends gnupg \
    && echo "deb http://ppa.launchpad.net/ondrej/php/ubuntu bionic main" > /etc/apt/sources.list.d/ondrej-php.list \
    && apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 4F4EA0AAE5267A6C \
    && apt-get update \    
    && apt-get -y --no-install-recommends install \
       curl \
       ca-certificates \
       unzip \
       tzdata \
       nano \
       # PHP Package
       php7.3 \
       php7.3-apcu \
       php7.3-bcmath \
       php7.3-cli \
       php7.3-common \
       php7.3-curl \
       php7.3-fileinfo \
       php7.3-fpm \
       php7.3-gd \
       php7.3-imap \
       php7.3-json \
       php7.3-mbstring \
       php7.3-memcached \
       php7.3-mongodb \
       php7.3-mysqlnd \
       php7.3-msgpack \
       php7.3-opcache \
       php7.3-redis \
       php7.3-simplexml \
       php7.3-soap \
       php7.3-xml \
       php7.3-zip \
       # Nginx
       nginx \
       # Supervisor
       supervisor \
       # Cron
       cron \
    # Enable Opcache CLI
    && echo 'opcache.enable_cli=1' >> /etc/php/7.3/mods-available/opcache.ini \
    # Link php-fpm binary and config
    && ln -s /usr/sbin/php-fpm7.3 /usr/sbin/php-fpm \
    && ln -s /etc/php/7.3/fpm/php-fpm.conf /etc/php/php-fpm.conf \
    # Install composer
    && curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer --1 \
    # Clear cache composer
    && composer clear-cache \
    # Remove unused folder nginx
    && rm -rf /etc/nginx/modules-available \
              /etc/nginx/modules-enabled \
              /etc/nginx/sites-available \
              /etc/nginx/sites-enabled \
              /etc/nginx/snippets \
              /etc/nginx/proxy_params \ 
              /etc/nginx/uwsgi_params \
              /var/log/nginx/* \
              /var/www/html/* \
    # Create folder supervisord
    && rm -rf /etc/supervisor \
    && mkdir -p /etc/supervisord.d \
    # Update timezone
    && cp -r /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime \
    # Remove user unused
    && userdel www-data \
    && userdel games \
    && userdel news \
    && userdel _apt \
    && userdel backup \
    && userdel list \
    && userdel gnats \
    && useradd -M -d /var/www/html -s /usr/sbin/nologin nginx \
    && mkdir -p /var/log/php-fpm \
    && chown nginx /var/log/php-fpm /var/log/nginx -R \
    # Add crontab Laravel
    && echo '* * * * * /usr/bin/php /var/www/html/artisan schedule:run >> /dev/null 2>&1' > /var/spool/cron/crontabs/nginx \
    && chmod 600 /var/spool/cron/crontabs/nginx \
    && chown nginx:crontab /var/spool/cron/crontabs/nginx \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* /usr/share/doc/* ~/.composer

ADD conf/nginx.conf /etc/nginx/nginx.conf
ADD conf/nginx-site.conf /etc/nginx/conf.d/vhost.conf

ADD conf/supervisord.conf /etc/supervisord.conf
ADD conf/nginx_php-fpm.conf /etc/supervisord.d/nginx_php-fpm.conf

ADD conf/www.conf /etc/php/7.3/fpm/pool.d/www.conf

ADD scripts/start.sh /start.sh
RUN chmod 700 /start.sh

WORKDIR /var/www/html

CMD ["/start.sh"]
