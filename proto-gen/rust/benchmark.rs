// @generated
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct ExecuteRequest {
    #[prost(message, repeated, tag="1")]
    pub tasks: ::prost::alloc::vec::Vec<Task>,
    /// for logging
    #[prost(string, tag="2")]
    pub group_id: ::prost::alloc::string::String,
    #[prost(string, tag="3")]
    pub contest_slug: ::prost::alloc::string::String,
    #[prost(enumeration="super::backend::Validator", tag="4")]
    pub validator: i32,
    /// sec
    #[prost(int32, tag="5")]
    pub time_limit_per_task: i32,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Task {
    #[prost(message, optional, tag="1")]
    pub request: ::core::option::Option<HttpRequest>,
    /// the number of threads for a task
    #[prost(int32, tag="6")]
    pub thread_num: i32,
    /// the count of attempting for a task
    #[prost(int32, tag="7")]
    pub attempt_count: i32,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct ExecuteResponse {
    #[prost(bool, tag="1")]
    pub ok: bool,
    /// if ok is false, this field is set
    #[prost(string, optional, tag="2")]
    pub error_message: ::core::option::Option<::prost::alloc::string::String>,
    /// in milliseconds
    #[prost(int64, tag="3")]
    pub time_elapsed: i64,
    #[prost(int32, tag="4")]
    pub total_requests: i32,
    #[prost(int32, tag="5")]
    pub requests_per_second: i32,
    #[prost(message, optional, tag="6")]
    pub task: ::core::option::Option<Task>,
    #[prost(enumeration="super::backend::Status", tag="7")]
    pub status: i32,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct HttpRequest {
    /// e.g.) <http://10.255.255.255/endpoint>
    #[prost(string, tag="1")]
    pub url: ::prost::alloc::string::String,
    #[prost(enumeration="HttpMethod", tag="2")]
    pub method: i32,
    #[prost(string, tag="4")]
    pub content_type: ::prost::alloc::string::String,
    #[prost(string, tag="5")]
    pub body: ::prost::alloc::string::String,
}
#[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
#[repr(i32)]
pub enum HttpMethod {
    Get = 0,
    Post = 1,
    Put = 2,
    Delete = 3,
}
impl HttpMethod {
    /// String value of the enum field names used in the ProtoBuf definition.
    ///
    /// The values are not transformed in any way and thus are considered stable
    /// (if the ProtoBuf definition does not change) and safe for programmatic use.
    pub fn as_str_name(&self) -> &'static str {
        match self {
            HttpMethod::Get => "GET",
            HttpMethod::Post => "POST",
            HttpMethod::Put => "PUT",
            HttpMethod::Delete => "DELETE",
        }
    }
    /// Creates an enum from field names used in the ProtoBuf definition.
    pub fn from_str_name(value: &str) -> ::core::option::Option<Self> {
        match value {
            "GET" => Some(Self::Get),
            "POST" => Some(Self::Post),
            "PUT" => Some(Self::Put),
            "DELETE" => Some(Self::Delete),
            _ => None,
        }
    }
}
include!("benchmark.tonic.rs");
// @@protoc_insertion_point(module)