.PHONY: all help translate test clean update compass collect rebuild

install:
	pip install -r requirements.txt

makemigrations:
	python manage.py makemigrations

migrate:
	python manage.py migrate

run:
	python manage.py runserver

test:
	python manage.py test