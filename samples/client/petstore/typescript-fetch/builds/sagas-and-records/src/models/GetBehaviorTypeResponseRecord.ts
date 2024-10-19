/* tslint:disable */
/* eslint-disable */
/**
 * OpenAPI Petstore
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import {ApiRecordUtils, knownRecordFactories, appFromJS, NormalizedRecordEntities} from "../runtimeSagasAndRecords";
import {getApiEntitiesState} from "../ApiEntitiesSelectors"
import {List, Record, RecordOf, Map} from 'immutable';
import {Schema, schema, NormalizedSchema} from "normalizr";
import {select, call} from "redux-saga/effects";

import {
    GetBehaviorTypeResponse,
} from './GetBehaviorTypeResponse';

import {
    BehaviorType,
} from './BehaviorType';
import {
    ResponseMeta,
} from './ResponseMeta';

import {
    ResponseMetaRecord,
    responseMetaRecordUtils
} from './ResponseMetaRecord';

export const GetBehaviorTypeResponseRecordProps = {
    recType: "GetBehaviorTypeResponseApiRecord" as "GetBehaviorTypeResponseApiRecord",
    meta: ResponseMetaRecord(),
    data: null as BehaviorType | null,
};

export type GetBehaviorTypeResponseRecordPropsType = typeof GetBehaviorTypeResponseRecordProps;
export const GetBehaviorTypeResponseRecord = Record(GetBehaviorTypeResponseRecordProps, GetBehaviorTypeResponseRecordProps.recType);
export type GetBehaviorTypeResponseRecord = RecordOf<GetBehaviorTypeResponseRecordPropsType>;

knownRecordFactories.set(GetBehaviorTypeResponseRecordProps.recType, GetBehaviorTypeResponseRecord);


class GetBehaviorTypeResponseRecordUtils extends ApiRecordUtils<GetBehaviorTypeResponse, GetBehaviorTypeResponseRecord> {
    public normalize(apiObject: GetBehaviorTypeResponse, asEntity?: boolean): GetBehaviorTypeResponse {
        (apiObject as any).recType = GetBehaviorTypeResponseRecordProps.recType;
        responseMetaRecordUtils.normalize(apiObject['meta']);
        return apiObject;
    }

    public toApi(record: GetBehaviorTypeResponseRecord): GetBehaviorTypeResponse {
        const apiObject = super.toApi(record);
        apiObject['meta'] = responseMetaRecordUtils.toApi(record['meta']);
        return apiObject;
    }

    public fromApiPassthrough(apiObject: GetBehaviorTypeResponse): BehaviorType {
        return apiObject.data!;
    }

    public fromApiPassthroughAsEntities(apiObject: GetBehaviorTypeResponse): NormalizedRecordEntities {
        console.log("entities revival not supported on this response");
        return {entities: {}, result: List<string>()};
    }
}

export const getBehaviorTypeResponseRecordUtils = new GetBehaviorTypeResponseRecordUtils();

