FROM python:latest

WORKDIR /app

COPY . .

RUN pip install --upgrade pip

RUN pip install -r requirements.txt

RUN touch banco.db

RUN python create_table.py

CMD ["python", "main.py"]