# alpine: light version
# rc: release candidate: bug fixing
From python:3.8-alpine

RUN pip install flask

WORKDIR /app

COPY . /app

EXPOSE 5000

CMD ["python", "main.py"]
