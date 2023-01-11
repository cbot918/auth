FROM python:3.10-alpine

RUN apk add --no-cache vim curl

WORKDIR /app/apitest

COPY requirements.txt .

RUN pip install pytest==7.2.0
RUN pip install -r requirements.txt