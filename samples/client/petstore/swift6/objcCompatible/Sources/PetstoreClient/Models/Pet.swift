//
// Pet.swift
//
// Generated by openapi-generator
// https://openapi-generator.tech
//

import Foundation

@objcMembers public class Pet: NSObject, Codable, JSONEncodable {

    public enum Status: String, Codable, CaseIterable {
        case available = "available"
        case pending = "pending"
        case sold = "sold"
    }    public var _id: Int64?
    public var _idNum: NSNumber? {
        get {
            return _id as NSNumber?
        }
    }
    public var category: Category?
    public var name: String
    public var photoUrls: [String]
    public var tags: [Tag]?
    /** pet status in the store */
    public var status: NullEncodable<Status>

    public init(_id: Int64? = nil, category: Category? = nil, name: String, photoUrls: [String], tags: [Tag]? = nil, status: NullEncodable<Status> = .encodeNull) {
        self._id = _id
        self.category = category
        self.name = name
        self.photoUrls = photoUrls
        self.tags = tags
        self.status = status
    }

    public enum CodingKeys: String, CodingKey, CaseIterable {
        case _id = "id"
        case category
        case name
        case photoUrls
        case tags
        case status
    }

    // Encodable protocol methods

    public func encode(to encoder: Encoder) throws {
        var container = encoder.container(keyedBy: CodingKeys.self)
        try container.encodeIfPresent(_id, forKey: ._id)
        try container.encodeIfPresent(category, forKey: .category)
        try container.encode(name, forKey: .name)
        try container.encode(photoUrls, forKey: .photoUrls)
        try container.encodeIfPresent(tags, forKey: .tags)
        switch status {
        case .encodeNothing: break
        case .encodeNull, .encodeValue: try container.encode(status, forKey: .status)
        }
    }
}

