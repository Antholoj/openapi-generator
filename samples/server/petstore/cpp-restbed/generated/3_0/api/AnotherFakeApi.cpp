/**
 * OpenAPI Petstore
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI-Generator 7.10.0-SNAPSHOT.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


#include <corvusoft/restbed/byte.hpp>
#include <corvusoft/restbed/string.hpp>
#include <corvusoft/restbed/settings.hpp>
#include <corvusoft/restbed/request.hpp>
#include <corvusoft/restbed/uri.hpp>
#include <boost/property_tree/ptree.hpp>
#include <boost/property_tree/json_parser.hpp>
#include <boost/lexical_cast.hpp>
#include <boost/algorithm/string.hpp>

#include "AnotherFakeApi.h"

namespace org {
namespace openapitools {
namespace server {
namespace api {

using namespace org::openapitools::server::model;

namespace {
[[maybe_unused]]
std::string selectPreferredContentType(const std::vector<std::string>& contentTypes) {
    if (contentTypes.size() == 0) {
        return "application/json";
    }

    if (contentTypes.size() == 1) {
        return contentTypes.at(0);
    }

    static const std::array<std::string, 2> preferredTypes = {"json", "xml"};
    for (const auto& preferredType: preferredTypes) {
        const auto ret = std::find_if(contentTypes.cbegin(),
        contentTypes.cend(),
        [preferredType](const std::string& str) {
            return str.find(preferredType) != std::string::npos;});
        if (ret != contentTypes.cend()) {
            return *ret;
        }
    }

    return contentTypes.at(0);
}
}

AnotherFakeApiException::AnotherFakeApiException(int status_code, std::string what)
  : m_status(status_code),
    m_what(what)
{

}
int AnotherFakeApiException::getStatus() const
{
    return m_status;
}
const char* AnotherFakeApiException::what() const noexcept
{
    return m_what.c_str();
}


template<class MODEL_T>
MODEL_T extractJsonModelBodyParam(const std::string& bodyContent)
{
    std::stringstream sstream(bodyContent);
    boost::property_tree::ptree pt;
    boost::property_tree::json_parser::read_json(sstream, pt);

    auto model = MODEL_T(pt);
    return model;
}

template<class MODEL_T>
std::vector<MODEL_T> extractJsonArrayBodyParam(const std::string& bodyContent)
{
    std::stringstream sstream(bodyContent);
    boost::property_tree::ptree pt;
    boost::property_tree::json_parser::read_json(sstream, pt);

    auto arrayRet = std::vector<MODEL_T>();
    for (const auto& child: pt) {
        arrayRet.emplace_back(MODEL_T(child.second));
    }
    return arrayRet;
}

template <class KEY_T, class VAL_T>
std::string convertMapResponse(const std::map<KEY_T, VAL_T>& map)
{
    boost::property_tree::ptree pt;
    for(const auto &kv: map) {
    pt.push_back(boost::property_tree::ptree::value_type(
        boost::lexical_cast<std::string>(kv.first),
        boost::property_tree::ptree(
        boost::lexical_cast<std::string>(kv.second))));
    }
    std::stringstream sstream;
    write_json(sstream, pt);
    std::string result = sstream.str();
    return result;
}

namespace AnotherFakeApiResources {
Another_fakeDummyResource::Another_fakeDummyResource(const std::string& context /* = "/v2" */)
{
	this->set_path(context + "/another-fake/dummy");
	this->set_method_handler("PATCH",
		std::bind(&Another_fakeDummyResource::handler_PATCH_internal, this,
			std::placeholders::_1));
}

std::pair<int, std::string> Another_fakeDummyResource::handleAnotherFakeApiException(const AnotherFakeApiException& e)
{
    return std::make_pair<int, std::string>(e.getStatus(), e.what());
}

std::pair<int, std::string> Another_fakeDummyResource::handleStdException(const std::exception& e)
{
    return std::make_pair<int, std::string>(500, e.what());
}

std::pair<int, std::string> Another_fakeDummyResource::handleUnspecifiedException()
{
    return std::make_pair<int, std::string>(500, "Unknown exception occurred");
}

void Another_fakeDummyResource::setResponseHeader(const std::shared_ptr<restbed::Session>& session, const std::string& header)
{
    session->set_header(header, "");
}

void Another_fakeDummyResource::returnResponse(const std::shared_ptr<restbed::Session>& session, const int status, const std::string& result, std::multimap<std::string, std::string>& responseHeaders)
{
    responseHeaders.insert(std::make_pair("Connection", "close"));
    session->close(status, result, responseHeaders);
}

void Another_fakeDummyResource::defaultSessionClose(const std::shared_ptr<restbed::Session>& session, const int status, const std::string& result)
{
    session->close(status, result, { {"Connection", "close"} });
}

void Another_fakeDummyResource::handler_PATCH_internal(const std::shared_ptr<restbed::Session> session)
{
    const auto request = session->get_request();
    // body params or form params here from the body content string
    std::string bodyContent = extractBodyContent(session);
    auto client = extractJsonModelBodyParam<Client>(bodyContent);
    
    int status_code = 500;
    Client resultObject = Client{};
    std::string result = "";
    
    try {
        std::tie(status_code, resultObject) =
            handler_PATCH(client);
    }
    catch(const AnotherFakeApiException& e) {
        std::tie(status_code, result) = handleAnotherFakeApiException(e);
    }
    catch(const std::exception& e) {
        std::tie(status_code, result) = handleStdException(e);
    }
    catch(...) {
        std::tie(status_code, result) = handleUnspecifiedException();
    }
    
    std::multimap< std::string, std::string > responseHeaders {};
    static const std::vector<std::string> contentTypes{
        "application/json",
    };
    static const std::string acceptTypes{
        "application/json, "
    };
    
    if (status_code == 200) {
        responseHeaders.insert(std::make_pair("Content-Type", selectPreferredContentType(contentTypes)));
        if (!acceptTypes.empty()) {
            responseHeaders.insert(std::make_pair("Accept", acceptTypes));
        }
    
        result = resultObject.toJsonString();
        returnResponse(session, 200, result.empty() ? "{}" : result, responseHeaders);
        return;
    }
    defaultSessionClose(session, status_code, result);
    
}


std::pair<int, Client> Another_fakeDummyResource::handler_PATCH(
        Client & client)
{
    return handler_PATCH_func(client);
}


std::string Another_fakeDummyResource::extractBodyContent(const std::shared_ptr<restbed::Session>& session) {
  const auto request = session->get_request();
  int content_length = request->get_header("Content-Length", 0);
  std::string bodyContent;
  session->fetch(content_length,
                 [&bodyContent](const std::shared_ptr<restbed::Session> session,
                                const restbed::Bytes &body) {
                   bodyContent = restbed::String::format(
                       "%.*s\n", (int)body.size(), body.data());
                 });
  return bodyContent;
}

std::string Another_fakeDummyResource::extractFormParamsFromBody(const std::string& paramName, const std::string& body) {
    const auto uri = restbed::Uri("urlencoded?" + body, true);
    const auto params = uri.get_query_parameters();
    const auto result = params.find(paramName);
    if (result != params.cend()) {
        return result->second;
    }
    return "";
}

} /* namespace AnotherFakeApiResources */

AnotherFakeApi::AnotherFakeApi(std::shared_ptr<restbed::Service> const& restbedService)
: m_service(restbedService)
{
}

AnotherFakeApi::~AnotherFakeApi() {}

std::shared_ptr<AnotherFakeApiResources::Another_fakeDummyResource> AnotherFakeApi::getAnother_fakeDummyResource() {
    if (!m_spAnother_fakeDummyResource) {
        setResource(std::make_shared<AnotherFakeApiResources::Another_fakeDummyResource>());
    }
    return m_spAnother_fakeDummyResource;
}
void AnotherFakeApi::setResource(std::shared_ptr<AnotherFakeApiResources::Another_fakeDummyResource> resource) {
    m_spAnother_fakeDummyResource = resource;
    m_service->publish(m_spAnother_fakeDummyResource);
}
void AnotherFakeApi::setAnotherFakeApiAnother_fakeDummyResource(std::shared_ptr<AnotherFakeApiResources::Another_fakeDummyResource> spAnother_fakeDummyResource) {
    m_spAnother_fakeDummyResource = spAnother_fakeDummyResource;
    m_service->publish(m_spAnother_fakeDummyResource);
}


void AnotherFakeApi::publishDefaultResources() {
    if (!m_spAnother_fakeDummyResource) {
        setResource(std::make_shared<AnotherFakeApiResources::Another_fakeDummyResource>());
    }
}

std::shared_ptr<restbed::Service> AnotherFakeApi::service() {
    return m_service;
}


}
}
}
}

