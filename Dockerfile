FROM python:3.8.6

#create folder
RUN mkdir /src
#set Directory
WORKDIR /src

ADD requirements.txt /src/
ADD entrypoint.sh /src/

ENV PYTHONUNBUFFERED=1
RUN pip install -r requirements.txt

ADD . /src/

USER root
RUN  chmod +x /src/entrypoint.sh
