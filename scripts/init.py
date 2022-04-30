import sys
import json

from urllib.error import HTTPError, URLError
from urllib.request import urlopen, Request

API_HOST = "http://localhost"
API_CONTACT = "/api/v1/contacts"


def get_url(argv):
    if len(argv) < 2:
        print("usage: {} <port_number>".format(argv[0]))
        sys.exit(1)

    host_port = argv[1]

    return "{}:{}{}".format(API_HOST, host_port, API_CONTACT)


def make_request(url, data=None, method="GET", headers=None):
    request = Request(
        url,
        data=data,
        method=method,
        headers=headers or {"Content-Type": "application/json"}
    )

    try:
        with urlopen(request, timeout=10) as response:
            return response.status, response.read()
    except HTTPError as error:
        return error.status, None
    except URLError as error:
        print(error.reason)
        return None, None
    except TimeoutError:
        print("Request timed out")
        return None, None


def encode_data(data):
    assert type(data) is dict
    data_json = json.dumps(data)
    return data_json.encode("utf-8")


def check_contact(contact):
    assert type(contact) is dict
    assert "id" in contact
    assert "name" in contact
    assert "email" in contact
