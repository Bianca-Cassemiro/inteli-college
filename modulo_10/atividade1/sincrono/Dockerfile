FROM python:latest

WORKDIR /app

COPY . .

RUN pip3 install --upgrade pip

RUN pip3 install -r requirements.txt

RUN touch sqlite.db

RUN python create_table.py

CMD ["python", "main.py"]