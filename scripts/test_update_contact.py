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

    contact_id = contact["id"]  # store the id of the newly created contact
    del contact["id"]
    assert contact == data_dict

    new_data_dict = {"name": "kurnia", "email": "kurnia@kurnia.com"}
    new_data = encode_data(new_data_dict)

    id_url = "{}/{}".format(url, contact_id)
    code, body = make_request(id_url, new_data, "PUT")
    assert code == 200

    body_dict = json.loads(body)
    check_contact(body_dict)

    del body_dict["id"]
    assert body_dict == new_data_dict

    code, body = make_request(id_url)
    assert code == 200

    body_dict = json.loads(body)
    check_contact(body_dict)

    del body_dict["id"]
    assert body_dict == new_data_dict

    print("all tests passed")
