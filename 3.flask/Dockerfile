#From ubuntu:20.04
FROM python:3.8-alpine

RUN pip install flask
#RUN pip install -r requirements.txt

WORKDIR /app
# ADD: copied after uncapsulating, COPY: copied itself
COPY app.py /app
COPY templates /app/templates

CMD ["python", "app.py"]

# formal, no need
EXPOSE 5000