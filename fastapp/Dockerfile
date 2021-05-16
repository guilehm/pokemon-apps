FROM python:3.8
ADD . /code
WORKDIR /code
COPY pyproject.toml /tmp/
RUN pip install poetry
RUN cd /tmp && poetry export -f requirements.txt
RUN pip install -r /tmp/requirements.txt
