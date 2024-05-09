FROM python:3.11

WORKDIR /code

COPY ./app.py /code/app.py
COPY ./noteAi /code/noteAi
COPY ./requirements.txt /code/requirements.txt

RUN pip install --no-cache-dir -r /code/requirements.txt

CMD ["fastapi", "run", "app.py", "--port", "8000"]
