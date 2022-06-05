.PHONY: all help translate test clean update compass collect rebuild

run:
	python manage.py runserver

test:
	python manage.py test