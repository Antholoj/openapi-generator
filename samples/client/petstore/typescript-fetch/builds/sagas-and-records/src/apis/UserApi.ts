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

import * as runtime from '../runtime';
import type {
  DefaultMetaOnlyResponse,
  User,
} from '../models/index';
import {
    DefaultMetaOnlyResponseFromJSON,
    DefaultMetaOnlyResponseToJSON,
    UserFromJSON,
    UserToJSON,
} from '../models/index';

export interface CreateUserRequest {
    body: User;
}

export interface CreateUsersWithArrayInputRequest {
    body: Array<User>;
}

export interface CreateUsersWithListInputRequest {
    body: Array<User>;
}

export interface DeleteUserRequest {
    username: string;
}

export interface GetUserByNameRequest {
    username: string;
}

export interface LoginUserRequest {
    username: string;
    password: string;
}

export interface UpdateUserRequest {
    username: string;
    body: User;
}

/**
 * 
 */
export class UserApi extends runtime.BaseAPI {

    /**
     * This can only be done by the logged in user.
     * Create user
     */
    async createUserRaw(requestParameters: CreateUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['body'] == null) {
            throw new runtime.RequiredError(
                'body',
                'Required parameter "body" was null or undefined when calling createUser().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/user`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: UserToJSON(requestParameters['body']),
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * This can only be done by the logged in user.
     * Create user
     */
    async createUser(body: User, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.createUserRaw({ body: body }, initOverrides);
    }

    /**
     * Creates list of users with given input array
     */
    async createUsersWithArrayInputRaw(requestParameters: CreateUsersWithArrayInputRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['body'] == null) {
            throw new runtime.RequiredError(
                'body',
                'Required parameter "body" was null or undefined when calling createUsersWithArrayInput().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/user/createWithArray`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: requestParameters['body']!.map(UserToJSON),
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Creates list of users with given input array
     */
    async createUsersWithArrayInput(body: Array<User>, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.createUsersWithArrayInputRaw({ body: body }, initOverrides);
    }

    /**
     * Creates list of users with given input array
     */
    async createUsersWithListInputRaw(requestParameters: CreateUsersWithListInputRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['body'] == null) {
            throw new runtime.RequiredError(
                'body',
                'Required parameter "body" was null or undefined when calling createUsersWithListInput().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/user/createWithList`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: requestParameters['body']!.map(UserToJSON),
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Creates list of users with given input array
     */
    async createUsersWithListInput(body: Array<User>, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.createUsersWithListInputRaw({ body: body }, initOverrides);
    }

    /**
     * This can only be done by the logged in user.
     * Delete user
     */
    async deleteUserRaw(requestParameters: DeleteUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['username'] == null) {
            throw new runtime.RequiredError(
                'username',
                'Required parameter "username" was null or undefined when calling deleteUser().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/user/{username}`.replace(`{${"username"}}`, encodeURIComponent(String(requestParameters['username']))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * This can only be done by the logged in user.
     * Delete user
     */
    async deleteUser(username: string, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.deleteUserRaw({ username: username }, initOverrides);
    }

    /**
     * Get user by user name
     */
    async getUserByNameRaw(requestParameters: GetUserByNameRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<User>> {
        if (requestParameters['username'] == null) {
            throw new runtime.RequiredError(
                'username',
                'Required parameter "username" was null or undefined when calling getUserByName().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/user/{username}`.replace(`{${"username"}}`, encodeURIComponent(String(requestParameters['username']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => UserFromJSON(jsonValue));
    }

    /**
     * Get user by user name
     */
    async getUserByName(username: string, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<User> {
        const response = await this.getUserByNameRaw({ username: username }, initOverrides);
        return await response.value();
    }

    /**
     * Logs user into the system
     */
    async loginUserRaw(requestParameters: LoginUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<string>> {
        if (requestParameters['username'] == null) {
            throw new runtime.RequiredError(
                'username',
                'Required parameter "username" was null or undefined when calling loginUser().'
            );
        }

        if (requestParameters['password'] == null) {
            throw new runtime.RequiredError(
                'password',
                'Required parameter "password" was null or undefined when calling loginUser().'
            );
        }

        const queryParameters: any = {};

        if (requestParameters['username'] != null) {
            queryParameters['username'] = requestParameters['username'];
        }

        if (requestParameters['password'] != null) {
            queryParameters['password'] = requestParameters['password'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/user/login`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        if (this.isJsonMime(response.headers.get('content-type'))) {
            return new runtime.JSONApiResponse<string>(response);
        } else {
            return new runtime.TextApiResponse(response) as any;
        }
    }

    /**
     * Logs user into the system
     */
    async loginUser(username: string, password: string, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<string> {
        const response = await this.loginUserRaw({ username: username, password: password }, initOverrides);
        return await response.value();
    }

    /**
     * Logs out current logged in user session
     */
    async logoutUserRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/user/logout`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Logs out current logged in user session
     */
    async logoutUser(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.logoutUserRaw(initOverrides);
    }

    /**
     * This can only be done by the logged in user.
     * Updated user
     */
    async updateUserRaw(requestParameters: UpdateUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<DefaultMetaOnlyResponse>> {
        if (requestParameters['username'] == null) {
            throw new runtime.RequiredError(
                'username',
                'Required parameter "username" was null or undefined when calling updateUser().'
            );
        }

        if (requestParameters['body'] == null) {
            throw new runtime.RequiredError(
                'body',
                'Required parameter "body" was null or undefined when calling updateUser().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/user/{username}`.replace(`{${"username"}}`, encodeURIComponent(String(requestParameters['username']))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: UserToJSON(requestParameters['body']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => DefaultMetaOnlyResponseFromJSON(jsonValue));
    }

    /**
     * This can only be done by the logged in user.
     * Updated user
     */
    async updateUser(username: string, body: User, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<DefaultMetaOnlyResponse> {
        const response = await this.updateUserRaw({ username: username, body: body }, initOverrides);
        return await response.value();
    }

}
