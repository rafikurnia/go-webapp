#!/usr/bin/env python3

import sys
import json

from random import randint
from init import check_contact, get_url, make_request


if __name__ == "__main__":
    url = get_url(sys.argv)

    # Test retrieval of all contacts
    code, body = make_request(url)
    assert code == 200

    contacts = json.loads(body)
    assert type(contacts) is list

    ids = []
    for contact in contacts:
        # Test retrieval of each contact
        check_contact(contact)

        id = contact["id"]
        ids.append(id)

        id_url = "{}/{}".format(url, id)
        code, body = make_request(id_url)

        assert code == 200
        body_dict = json.loads(body)

        assert type(body_dict) is dict
        assert body_dict == contact

    # Find non-existence contact ID
    randId = 1
    while randId in ids:
        randId = randint(1, 1000)

    # Test retrieval of non-existence contact
    id_url = "{}/{}".format(url, randId)
    code, _ = make_request(id_url)
    assert code == 404

    print("all tests passed")
