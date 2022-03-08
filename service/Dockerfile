FROM php:8.0-fpm

# Install dependencies
RUN apt-get update && apt-get install -y \
    build-essential \
    libpng-dev \
    libjpeg62-turbo-dev \
    libfreetype6-dev \
    locales \
    jpegoptim optipng pngquant gifsicle \
    vim \
    unzip \
    git \
    curl
RUN apt-get install libonig-dev
# Clear cache
RUN apt-get clean && rm -rf /var/lib/apt/lists/*

# Install extensions
RUN docker-php-ext-install pdo_mysql mbstring exif pcntl \
    && docker-php-ext-install mbstring mysqli opcache pdo_mysql \
    && apt-get purge -y --auto-remove $buildDeps 
    
RUN mkdir -p /var/www/html/
RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer

ADD . /var/www/html/
# RUN composer update
RUN composer install --no-scripts --no-autoloader

RUN composer dump-autoload --optimize