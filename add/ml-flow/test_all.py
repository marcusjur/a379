# tests/test_autocompletion.py
import pytest
import requests
import ctypes
import json
import os

os.environ['DYLD_LIBRARY_PATH'] = os.getcwd()
if 'DOCKER_CONTAINER_ID' in os.environ:
    libsession = ctypes.CDLL('./libsession.so')
else:
    libsession = ctypes.CDLL('../go-flow/session/libsession.so')
libsession.calculateSessionKey.restype = ctypes.c_char_p
session_key = libsession.calculateSessionKey().decode('utf-8')


@pytest.mark.parametrize('session_key', [session_key])
def test_autocompletion_endpoint(session_key):
    url = 'http://localhost:8000/api/v1/autocompletion'

    headers = {'Authorization': f'{session_key}', 'Content-Type': 'application/json'}
    print(headers)

    data = {'text': 'I would like a pizza with'}

    response = requests.get(url, headers=headers, json=data)
    assert response.status_code == 200
    resp = json.loads(response.json())
    print(resp)

    # Assert that the response was successful
    assert isinstance(resp['generated_text'], str)
    assert len(resp['generated_text']) > 3
