FROM python:3.12-slim
WORKDIR /app-backend
RUN pip install poetry

COPY pyproject.toml ./
COPY add/go-flow/session/libsession.so ./
COPY add/go-flow/session/libsession.h ./

RUN poetry install

COPY add/ml-flow/autocompletion.py ./

EXPOSE 8000

CMD ["poetry", "run", "python", "autocompletion.py"]
