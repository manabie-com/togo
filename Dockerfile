FROM python:3.8.12

WORKDIR /app
ADD . /app

RUN pip install --upgrade pip

RUN pip install -r requirements.txt

EXPOSE 8000

CMD ["python", "app.py", "0.0.0.0:8000"]