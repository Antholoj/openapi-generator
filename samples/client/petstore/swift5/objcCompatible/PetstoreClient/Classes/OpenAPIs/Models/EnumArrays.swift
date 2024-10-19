//
// EnumArrays.swift
//
// Generated by openapi-generator
// https://openapi-generator.tech
//

import Foundation
#if canImport(AnyCodable)
import AnyCodable
#endif

@objcMembers public class EnumArrays: NSObject, Codable, JSONEncodable {

    public enum JustSymbol: String, Codable, CaseIterable {
        case greaterThanOrEqualTo = ">="
        case dollar = "$"
    }    public enum ArrayEnum: String, Codable, CaseIterable {
        case fish = "fish"
        case crab = "crab"
    }    public var justSymbol: JustSymbol?
    public var arrayEnum: [ArrayEnum]?

    public init(justSymbol: JustSymbol? = nil, arrayEnum: [ArrayEnum]? = nil) {
        self.justSymbol = justSymbol
        self.arrayEnum = arrayEnum
    }

    public enum CodingKeys: String, CodingKey, CaseIterable {
        case justSymbol = "just_symbol"
        case arrayEnum = "array_enum"
    }

    // Encodable protocol methods

    public func encode(to encoder: Encoder) throws {
        var container = encoder.container(keyedBy: CodingKeys.self)
        try container.encodeIfPresent(justSymbol, forKey: .justSymbol)
        try container.encodeIfPresent(arrayEnum, forKey: .arrayEnum)
    }
}

