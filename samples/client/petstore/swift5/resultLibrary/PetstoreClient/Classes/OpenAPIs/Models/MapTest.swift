//
// MapTest.swift
//
// Generated by openapi-generator
// https://openapi-generator.tech
//

import Foundation
#if canImport(AnyCodable)
import AnyCodable
#endif

internal struct MapTest: Codable, JSONEncodable {

    internal enum MapOfEnumString: String, Codable, CaseIterable, CaseIterableDefaultsLast {
        case upper = "UPPER"
        case lower = "lower"
        case unknownDefaultOpenApi = "unknown_default_open_api"
    }    internal private(set) var mapMapOfString: [String: [String: String]]?
    internal private(set) var mapOfEnumString: [String: String]?
    internal private(set) var directMap: [String: Bool]?
    internal private(set) var indirectMap: StringBooleanMap?

    internal init(mapMapOfString: [String: [String: String]]? = nil, mapOfEnumString: [String: String]? = nil, directMap: [String: Bool]? = nil, indirectMap: StringBooleanMap? = nil) {
        self.mapMapOfString = mapMapOfString
        self.mapOfEnumString = mapOfEnumString
        self.directMap = directMap
        self.indirectMap = indirectMap
    }

    internal enum CodingKeys: String, CodingKey, CaseIterable {
        case mapMapOfString = "map_map_of_string"
        case mapOfEnumString = "map_of_enum_string"
        case directMap = "direct_map"
        case indirectMap = "indirect_map"
    }

    // Encodable protocol methods

    internal func encode(to encoder: Encoder) throws {
        var container = encoder.container(keyedBy: CodingKeys.self)
        try container.encodeIfPresent(mapMapOfString, forKey: .mapMapOfString)
        try container.encodeIfPresent(mapOfEnumString, forKey: .mapOfEnumString)
        try container.encodeIfPresent(directMap, forKey: .directMap)
        try container.encodeIfPresent(indirectMap, forKey: .indirectMap)
    }
}

