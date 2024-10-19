# coding: utf-8

"""
    OpenAPI Petstore

    This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\

    The version of the OpenAPI document: 1.0.0
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501

from __future__ import annotations
import pprint
import re  # noqa: F401
import json


from typing import List, Optional
from pydantic import BaseModel, Field, StrictStr, conlist

class CircularAllOfRef(BaseModel):
    """
    CircularAllOfRef
    """
    name: Optional[StrictStr] = Field(default=None, alias="_name")
    second_circular_all_of_ref: Optional[conlist(SecondCircularAllOfRef)] = Field(default=None, alias="secondCircularAllOfRef")
    __properties = ["_name", "secondCircularAllOfRef"]

    class Config:
        """Pydantic configuration"""
        allow_population_by_field_name = True
        validate_assignment = True

    def to_str(self) -> str:
        """Returns the string representation of the model using alias"""
        return pprint.pformat(self.dict(by_alias=True))

    def to_json(self) -> str:
        """Returns the JSON representation of the model using alias"""
        return json.dumps(self.to_dict())

    @classmethod
    def from_json(cls, json_str: str) -> CircularAllOfRef:
        """Create an instance of CircularAllOfRef from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                          },
                          exclude_none=True)
        # override the default output from pydantic by calling `to_dict()` of each item in second_circular_all_of_ref (list)
        _items = []
        if self.second_circular_all_of_ref:
            for _item in self.second_circular_all_of_ref:
                if _item:
                    _items.append(_item.to_dict())
            _dict['secondCircularAllOfRef'] = _items
        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> CircularAllOfRef:
        """Create an instance of CircularAllOfRef from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return CircularAllOfRef.parse_obj(obj)

        _obj = CircularAllOfRef.parse_obj({
            "name": obj.get("_name"),
            "second_circular_all_of_ref": [SecondCircularAllOfRef.from_dict(_item) for _item in obj.get("secondCircularAllOfRef")] if obj.get("secondCircularAllOfRef") is not None else None
        })
        return _obj

from petstore_api.models.second_circular_all_of_ref import SecondCircularAllOfRef
CircularAllOfRef.update_forward_refs()

