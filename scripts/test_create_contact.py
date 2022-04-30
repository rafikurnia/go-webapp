#!/usr/bin/env python3

import sys
import json

from init import check_contact, get_url, make_request, encode_data


if __name__ == "__main__":
    url = get_url(sys.argv)

    data_dict = {"name": "rafi", "email": "rafi@rafi.com"}
    data = encode_data(data_dict)

    code, body = make_request(url, data, "POST")
    assert code == 201

    contact = json.loads(body)
    check_contact(contact)

    del contact["id"]
    assert contact == data_dict

    print("all tests passed")
