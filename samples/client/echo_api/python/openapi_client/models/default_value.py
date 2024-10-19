# coding: utf-8

"""
    Echo Server API

    Echo Server API

    The version of the OpenAPI document: 0.1.0
    Contact: team@openapitools.org
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501

from __future__ import annotations
import pprint
import re  # noqa: F401
import json

from pydantic import BaseModel, ConfigDict, StrictInt, StrictStr, field_validator
from typing import Any, ClassVar, Dict, List, Optional
from openapi_client.models.string_enum_ref import StringEnumRef
from typing import Optional, Set
from typing_extensions import Self

class DefaultValue(BaseModel):
    """
    to test the default value of properties
    """ # noqa: E501
    array_string_enum_ref_default: Optional[List[StringEnumRef]] = None
    array_string_enum_default: Optional[List[StrictStr]] = None
    array_string_default: Optional[List[StrictStr]] = None
    array_integer_default: Optional[List[StrictInt]] = None
    array_string: Optional[List[StrictStr]] = None
    array_string_nullable: Optional[List[StrictStr]] = None
    array_string_extension_nullable: Optional[List[StrictStr]] = None
    string_nullable: Optional[StrictStr] = None
    __properties: ClassVar[List[str]] = ["array_string_enum_ref_default", "array_string_enum_default", "array_string_default", "array_integer_default", "array_string", "array_string_nullable", "array_string_extension_nullable", "string_nullable"]

    @field_validator('array_string_enum_default')
    def array_string_enum_default_validate_enum(cls, value):
        """Validates the enum"""
        if value is None:
            return value

        for i in value:
            if i not in set(['success', 'failure', 'unclassified']):
                raise ValueError("each list item must be one of ('success', 'failure', 'unclassified')")
        return value

    model_config = ConfigDict(
        populate_by_name=True,
        validate_assignment=True,
        protected_namespaces=(),
    )


    def to_str(self) -> str:
        """Returns the string representation of the model using alias"""
        return pprint.pformat(self.model_dump(by_alias=True))

    def to_json(self) -> str:
        """Returns the JSON representation of the model using alias"""
        # TODO: pydantic v2: use .model_dump_json(by_alias=True, exclude_unset=True) instead
        return json.dumps(self.to_dict())

    @classmethod
    def from_json(cls, json_str: str) -> Optional[Self]:
        """Create an instance of DefaultValue from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self) -> Dict[str, Any]:
        """Return the dictionary representation of the model using alias.

        This has the following differences from calling pydantic's
        `self.model_dump(by_alias=True)`:

        * `None` is only added to the output dict for nullable fields that
          were set at model initialization. Other fields with value `None`
          are ignored.
        """
        excluded_fields: Set[str] = set([
        ])

        _dict = self.model_dump(
            by_alias=True,
            exclude=excluded_fields,
            exclude_none=True,
        )
        # set to None if array_string_nullable (nullable) is None
        # and model_fields_set contains the field
        if self.array_string_nullable is None and "array_string_nullable" in self.model_fields_set:
            _dict['array_string_nullable'] = None

        # set to None if array_string_extension_nullable (nullable) is None
        # and model_fields_set contains the field
        if self.array_string_extension_nullable is None and "array_string_extension_nullable" in self.model_fields_set:
            _dict['array_string_extension_nullable'] = None

        # set to None if string_nullable (nullable) is None
        # and model_fields_set contains the field
        if self.string_nullable is None and "string_nullable" in self.model_fields_set:
            _dict['string_nullable'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of DefaultValue from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "array_string_enum_ref_default": obj.get("array_string_enum_ref_default"),
            "array_string_enum_default": obj.get("array_string_enum_default"),
            "array_string_default": obj.get("array_string_default"),
            "array_integer_default": obj.get("array_integer_default"),
            "array_string": obj.get("array_string"),
            "array_string_nullable": obj.get("array_string_nullable"),
            "array_string_extension_nullable": obj.get("array_string_extension_nullable"),
            "string_nullable": obj.get("string_nullable")
        })
        return _obj


