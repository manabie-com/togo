#!/bin/bash

# Generate random password for supervisord
SUPERVISORD_PASS=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w ${1:-16} | head -n 1)
sed -i s/^password=.*$/password=${SUPERVISORD_PASS}/g /etc/supervisord.conf

echo "$(date) - Set PHP variables"
# Increase the memory_limit
if [ ! -z "$PHP_MEM_LIMIT" ]; then
  sed -i "s/memory_limit = 128M/memory_limit = ${PHP_MEM_LIMIT}M/g" ${php_ini_file}
fi

# Increase the upload_max_filesize
if [ ! -z "$PHP_UPLOAD_MAX_FILESIZE" ]; then
  sed -i "s/upload_max_filesize = 2M/upload_max_filesize = ${PHP_UPLOAD_MAX_FILESIZE}M/g" ${php_ini_file}
  sed -i "s/post_max_size = 8M/post_max_size = ${PHP_UPLOAD_MAX_FILESIZE}M/g" ${php_ini_file}
  sed -i "s/client_max_body_size 2M/client_max_body_size ${PHP_UPLOAD_MAX_FILESIZE}M/g" /etc/nginx/nginx.conf
fi

echo "$(date) - Run custom script"
# Run custom scripts
if [[ "$RUN_SCRIPTS" == "1" ]] ; then
  if [ -d "/var/www/html/scripts/" ]; then
    # make scripts executable incase they aren't
    chmod -Rf 750 /var/www/html/scripts/*
    # run scripts in number order
    for i in `ls /var/www/html/scripts/`; do /var/www/html/scripts/$i ; done
  else
    echo "Can't find script directory"
  fi
fi

echo "$(date) Start supervisord and services"
# Start supervisord and services
exec /usr/bin/supervisord -n -c /etc/supervisord.conf