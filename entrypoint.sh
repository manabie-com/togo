#!/bin/bash

set -e
echo "Migrating the database before starting the server"
python manage.py makemigrations
python manage.py migrate
python manage.py collectstatic --noinput

echo "runserver"
gunicorn togo.wsgi:application --bind 0.0.0.0:8000