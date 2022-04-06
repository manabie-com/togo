FROM python:3.7-alpine
ENV PYTHONUNBUFFERED 1

COPY requirements requirements
COPY requirements.txt requirements.txt

RUN apk update && \
    apk add --no-cache --virtual build-deps gcc musl-dev python3-dev postgresql-dev
RUN apk add libpq

RUN pip install --trusted-host pypi.python.org -r requirements.txt

RUN apk del --no-cache build-deps gcc musl-dev

COPY . /app
WORKDIR /app

EXPOSE 5000

ENTRYPOINT [ "gunicorn", "-b", ":5000", "--log-level", "INFO", "manage:app" ]
