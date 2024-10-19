//
// EnumTest.swift
//
// Generated by openapi-generator
// https://openapi-generator.tech
//

import Foundation

@available(*, deprecated, renamed: "PetstoreClientAPI.EnumTest")
public typealias EnumTest = PetstoreClientAPI.EnumTest

extension PetstoreClientAPI {

public final class EnumTest: Codable, JSONEncodable, Hashable {

    public enum EnumString: String, Codable, CaseIterable {
        case upper = "UPPER"
        case lower = "lower"
        case empty = ""
    }    public enum EnumStringRequired: String, Codable, CaseIterable {
        case upper = "UPPER"
        case lower = "lower"
        case empty = ""
    }    public enum EnumInteger: Int, Codable, CaseIterable {
        case _1 = 1
        case number1 = -1
    }    public enum EnumNumber: Double, Codable, CaseIterable {
        case _11 = 1.1
        case number12 = -1.2
    }    public var enumString: EnumString?
    public var enumStringRequired: EnumStringRequired
    public var enumInteger: EnumInteger?
    public var enumNumber: EnumNumber?
    public var outerEnum: OuterEnum?

    public init(enumString: EnumString? = nil, enumStringRequired: EnumStringRequired, enumInteger: EnumInteger? = nil, enumNumber: EnumNumber? = nil, outerEnum: OuterEnum? = nil) {
        self.enumString = enumString
        self.enumStringRequired = enumStringRequired
        self.enumInteger = enumInteger
        self.enumNumber = enumNumber
        self.outerEnum = outerEnum
    }

    public enum CodingKeys: String, CodingKey, CaseIterable {
        case enumString = "enum_string"
        case enumStringRequired = "enum_string_required"
        case enumInteger = "enum_integer"
        case enumNumber = "enum_number"
        case outerEnum
    }

    // Encodable protocol methods

    public func encode(to encoder: Encoder) throws {
        var container = encoder.container(keyedBy: CodingKeys.self)
        try container.encodeIfPresent(enumString, forKey: .enumString)
        try container.encode(enumStringRequired, forKey: .enumStringRequired)
        try container.encodeIfPresent(enumInteger, forKey: .enumInteger)
        try container.encodeIfPresent(enumNumber, forKey: .enumNumber)
        try container.encodeIfPresent(outerEnum, forKey: .outerEnum)
    }

    public static func == (lhs: EnumTest, rhs: EnumTest) -> Bool {
        lhs.enumString == rhs.enumString &&
        lhs.enumStringRequired == rhs.enumStringRequired &&
        lhs.enumInteger == rhs.enumInteger &&
        lhs.enumNumber == rhs.enumNumber &&
        lhs.outerEnum == rhs.outerEnum
        
    }

    public func hash(into hasher: inout Hasher) {
        hasher.combine(enumString?.hashValue)
        hasher.combine(enumStringRequired.hashValue)
        hasher.combine(enumInteger?.hashValue)
        hasher.combine(enumNumber?.hashValue)
        hasher.combine(outerEnum?.hashValue)
        
    }
}

}
