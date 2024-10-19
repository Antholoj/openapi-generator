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



from pydantic import BaseModel, Field, StrictStr, validator

class Feeding(BaseModel):
    """
    Feeding
    """
    task_name: StrictStr = Field(...)
    function_name: StrictStr = Field(...)
    content: StrictStr = Field(...)
    __properties = ["task_name", "function_name", "content"]

    @validator('task_name')
    def task_name_validate_enum(cls, value):
        """Validates the enum"""
        if value not in ('cleaning'):
            raise ValueError("must be one of enum values ('cleaning')")
        return value

    @validator('function_name')
    def function_name_validate_enum(cls, value):
        """Validates the enum"""
        if value not in ('care_nourish'):
            raise ValueError("must be one of enum values ('care_nourish')")
        return value

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
    def from_json(cls, json_str: str) -> Feeding:
        """Create an instance of Feeding from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                          },
                          exclude_none=True)
        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> Feeding:
        """Create an instance of Feeding from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return Feeding.parse_obj(obj)

        _obj = Feeding.parse_obj({
            "task_name": obj.get("task_name"),
            "function_name": obj.get("function_name"),
            "content": obj.get("content")
        })
        return _obj


