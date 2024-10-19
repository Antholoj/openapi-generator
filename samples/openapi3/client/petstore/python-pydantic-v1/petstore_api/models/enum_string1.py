# coding: utf-8

"""
    OpenAPI Petstore

    This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\

    The version of the OpenAPI document: 1.0.0
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501

import json
import pprint
import re  # noqa: F401
from aenum import Enum, no_arg





class EnumString1(str, Enum):
    """
    EnumString1
    """

    """
    allowed enum values
    """
    A = 'a'
    B = 'b'

    @classmethod
    def from_json(cls, json_str: str) -> EnumString1:
        """Create an instance of EnumString1 from a JSON string"""
        return EnumString1(json.loads(json_str))

