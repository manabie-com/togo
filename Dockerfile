FROM python:3.7
ENV PYTHONDONTWRITEBYTECODE=1
ENV PYTHONUNBUFFERED=1

WORKDIR /usr/app/togo
COPY requirements.txt /usr/app/togo/
RUN python -m pip install -U pip
RUN pip install -r requirements.txt

COPY . /usr/app/togo/
