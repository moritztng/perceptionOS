# syntax=docker/dockerfile:1

FROM python:3.9-bullseye

WORKDIR /app

COPY requirements.txt .
RUN pip3.9 install -r requirements.txt

COPY perception perception
COPY docker/camera/app.py .

CMD [ "python3", "-m" , "flask", "run", "--host=0.0.0.0"]
