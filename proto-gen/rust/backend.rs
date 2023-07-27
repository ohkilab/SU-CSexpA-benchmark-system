// @generated
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Contest {
    #[prost(int32, tag="1")]
    pub id: i32,
    #[prost(string, tag="2")]
    pub title: ::prost::alloc::string::String,
    #[prost(message, optional, tag="4")]
    pub start_at: ::core::option::Option<::prost_types::Timestamp>,
    #[prost(message, optional, tag="5")]
    pub end_at: ::core::option::Option<::prost_types::Timestamp>,
    #[prost(int32, tag="6")]
    pub submit_limit: i32,
    #[prost(string, tag="8")]
    pub slug: ::prost::alloc::string::String,
    #[prost(enumeration="TagSelectionLogicType", tag="9")]
    pub tag_selection_logic: i32,
    #[prost(enumeration="Validator", tag="10")]
    pub validator: i32,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct TagSelectionLogicManual {
    #[prost(enumeration="TagSelectionLogicType", tag="1")]
    pub r#type: i32,
    /// tags_list\[i\] .. used if the attempt count is i+1
    #[prost(message, repeated, tag="2")]
    pub tags_list: ::prost::alloc::vec::Vec<Tags>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct TagSelectionLogicAuto {
    #[prost(enumeration="TagSelectionLogicType", tag="1")]
    pub r#type: i32,
    #[prost(message, optional, tag="2")]
    pub tags: ::core::option::Option<Tags>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Tags {
    #[prost(string, repeated, tag="1")]
    pub tags: ::prost::alloc::vec::Vec<::prost::alloc::string::String>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Group {
    #[prost(string, tag="1")]
    pub id: ::prost::alloc::string::String,
    #[prost(enumeration="Role", tag="4")]
    pub role: i32,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Submit {
    #[prost(int32, tag="1")]
    pub id: i32,
    #[prost(string, tag="2")]
    pub group_name: ::prost::alloc::string::String,
    #[prost(int32, tag="4")]
    pub score: i32,
    #[prost(enumeration="Language", tag="5")]
    pub language: i32,
    #[prost(message, optional, tag="6")]
    pub submited_at: ::core::option::Option<::prost_types::Timestamp>,
    /// it this field is not null, this submit is completed
    #[prost(message, optional, tag="7")]
    pub completed_at: ::core::option::Option<::prost_types::Timestamp>,
    #[prost(message, repeated, tag="8")]
    pub task_results: ::prost::alloc::vec::Vec<TaskResult>,
    #[prost(enumeration="Status", tag="9")]
    pub status: i32,
    /// if the connection error occurs, then this field is filled
    #[prost(string, optional, tag="10")]
    pub error_message: ::core::option::Option<::prost::alloc::string::String>,
    #[prost(int32, tag="11")]
    pub tag_count: i32,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct TaskResult {
    #[prost(int32, tag="1")]
    pub id: i32,
    #[prost(int32, tag="2")]
    pub request_per_sec: i32,
    #[prost(string, tag="3")]
    pub url: ::prost::alloc::string::String,
    #[prost(string, tag="4")]
    pub method: ::prost::alloc::string::String,
    #[prost(string, tag="5")]
    pub request_content_type: ::prost::alloc::string::String,
    #[prost(string, optional, tag="6")]
    pub request_body: ::core::option::Option<::prost::alloc::string::String>,
    #[prost(string, tag="7")]
    pub response_code: ::prost::alloc::string::String,
    #[prost(string, tag="8")]
    pub response_content_type: ::prost::alloc::string::String,
    #[prost(string, tag="9")]
    pub response_body: ::prost::alloc::string::String,
    #[prost(int32, tag="10")]
    pub thread_num: i32,
    #[prost(int32, tag="11")]
    pub attempt_count: i32,
    #[prost(int32, tag="12")]
    pub attempt_time: i32,
    #[prost(message, optional, tag="13")]
    pub created_at: ::core::option::Option<::prost_types::Timestamp>,
    #[prost(message, optional, tag="14")]
    pub deleted_at: ::core::option::Option<::prost_types::Timestamp>,
    #[prost(string, optional, tag="15")]
    pub error_message: ::core::option::Option<::prost::alloc::string::String>,
    #[prost(enumeration="Status", tag="16")]
    pub status: i32,
}
#[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
#[repr(i32)]
pub enum TagSelectionLogicType {
    Auto = 0,
    Manual = 1,
}
impl TagSelectionLogicType {
    /// String value of the enum field names used in the ProtoBuf definition.
    ///
    /// The values are not transformed in any way and thus are considered stable
    /// (if the ProtoBuf definition does not change) and safe for programmatic use.
    pub fn as_str_name(&self) -> &'static str {
        match self {
            TagSelectionLogicType::Auto => "AUTO",
            TagSelectionLogicType::Manual => "MANUAL",
        }
    }
    /// Creates an enum from field names used in the ProtoBuf definition.
    pub fn from_str_name(value: &str) -> ::core::option::Option<Self> {
        match value {
            "AUTO" => Some(Self::Auto),
            "MANUAL" => Some(Self::Manual),
            _ => None,
        }
    }
}
#[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
#[repr(i32)]
pub enum Status {
    /// waiting for benchmark
    Waiting = 0,
    /// in progress
    InProgress = 1,
    /// benchmark succeeded
    Success = 2,
    /// failed to connect
    ConnectionFailed = 3,
    /// validation error
    ValidationError = 4,
    /// backend error
    InternalError = 5,
    /// timeout
    Timeout = 6,
}
impl Status {
    /// String value of the enum field names used in the ProtoBuf definition.
    ///
    /// The values are not transformed in any way and thus are considered stable
    /// (if the ProtoBuf definition does not change) and safe for programmatic use.
    pub fn as_str_name(&self) -> &'static str {
        match self {
            Status::Waiting => "WAITING",
            Status::InProgress => "IN_PROGRESS",
            Status::Success => "SUCCESS",
            Status::ConnectionFailed => "CONNECTION_FAILED",
            Status::ValidationError => "VALIDATION_ERROR",
            Status::InternalError => "INTERNAL_ERROR",
            Status::Timeout => "TIMEOUT",
        }
    }
    /// Creates an enum from field names used in the ProtoBuf definition.
    pub fn from_str_name(value: &str) -> ::core::option::Option<Self> {
        match value {
            "WAITING" => Some(Self::Waiting),
            "IN_PROGRESS" => Some(Self::InProgress),
            "SUCCESS" => Some(Self::Success),
            "CONNECTION_FAILED" => Some(Self::ConnectionFailed),
            "VALIDATION_ERROR" => Some(Self::ValidationError),
            "INTERNAL_ERROR" => Some(Self::InternalError),
            "TIMEOUT" => Some(Self::Timeout),
            _ => None,
        }
    }
}
#[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
#[repr(i32)]
pub enum Language {
    Php = 0,
    Go = 1,
    Rust = 2,
    Javascript = 3,
    Csharp = 4,
    Cpp = 5,
    Ruby = 6,
    Python = 7,
}
impl Language {
    /// String value of the enum field names used in the ProtoBuf definition.
    ///
    /// The values are not transformed in any way and thus are considered stable
    /// (if the ProtoBuf definition does not change) and safe for programmatic use.
    pub fn as_str_name(&self) -> &'static str {
        match self {
            Language::Php => "PHP",
            Language::Go => "GO",
            Language::Rust => "RUST",
            Language::Javascript => "JAVASCRIPT",
            Language::Csharp => "CSHARP",
            Language::Cpp => "CPP",
            Language::Ruby => "RUBY",
            Language::Python => "PYTHON",
        }
    }
    /// Creates an enum from field names used in the ProtoBuf definition.
    pub fn from_str_name(value: &str) -> ::core::option::Option<Self> {
        match value {
            "PHP" => Some(Self::Php),
            "GO" => Some(Self::Go),
            "RUST" => Some(Self::Rust),
            "JAVASCRIPT" => Some(Self::Javascript),
            "CSHARP" => Some(Self::Csharp),
            "CPP" => Some(Self::Cpp),
            "RUBY" => Some(Self::Ruby),
            "PYTHON" => Some(Self::Python),
            _ => None,
        }
    }
}
#[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
#[repr(i32)]
pub enum Role {
    Contestant = 0,
    Guest = 1,
}
impl Role {
    /// String value of the enum field names used in the ProtoBuf definition.
    ///
    /// The values are not transformed in any way and thus are considered stable
    /// (if the ProtoBuf definition does not change) and safe for programmatic use.
    pub fn as_str_name(&self) -> &'static str {
        match self {
            Role::Contestant => "CONTESTANT",
            Role::Guest => "GUEST",
        }
    }
    /// Creates an enum from field names used in the ProtoBuf definition.
    pub fn from_str_name(value: &str) -> ::core::option::Option<Self> {
        match value {
            "CONTESTANT" => Some(Self::Contestant),
            "GUEST" => Some(Self::Guest),
            _ => None,
        }
    }
}
/// 運用的に難があるけど仕方ない・・
/// DB だけでここら辺をやるとしたら、AtCoder のスペシャルジャッジみたいに
/// シングルの Go や C++ で書かれた validator を download & compile して
/// request と response を渡してチェックしてもらうとかの形にしないといけない気がする
#[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
#[repr(i32)]
pub enum Validator {
    /// 2022年度
    V2022 = 0,
    /// 2023年度 予選
    V2023 = 1,
}
impl Validator {
    /// String value of the enum field names used in the ProtoBuf definition.
    ///
    /// The values are not transformed in any way and thus are considered stable
    /// (if the ProtoBuf definition does not change) and safe for programmatic use.
    pub fn as_str_name(&self) -> &'static str {
        match self {
            Validator::V2022 => "V2022",
            Validator::V2023 => "V2023",
        }
    }
    /// Creates an enum from field names used in the ProtoBuf definition.
    pub fn from_str_name(value: &str) -> ::core::option::Option<Self> {
        match value {
            "V2022" => Some(Self::V2022),
            "V2023" => Some(Self::V2023),
            _ => None,
        }
    }
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct PostLoginRequest {
    /// it means group_id
    #[prost(string, tag="1")]
    pub id: ::prost::alloc::string::String,
    #[prost(string, tag="2")]
    pub password: ::prost::alloc::string::String,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct PostLoginResponse {
    #[prost(message, optional, tag="1")]
    pub group: ::core::option::Option<Group>,
    #[prost(string, tag="2")]
    pub token: ::prost::alloc::string::String,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct PostSubmitRequest {
    #[prost(string, tag="1")]
    pub url: ::prost::alloc::string::String,
    #[prost(string, tag="2")]
    pub contest_slug: ::prost::alloc::string::String,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct PostSubmitResponse {
    #[prost(int32, tag="1")]
    pub id: i32,
    #[prost(string, tag="2")]
    pub url: ::prost::alloc::string::String,
    #[prost(string, tag="3")]
    pub contest_slug: ::prost::alloc::string::String,
    #[prost(message, optional, tag="4")]
    pub submited_at: ::core::option::Option<::prost_types::Timestamp>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct GetSubmitRequest {
    #[prost(int32, tag="1")]
    pub submit_id: i32,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct GetSubmitResponse {
    #[prost(message, optional, tag="1")]
    pub submit: ::core::option::Option<Submit>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct ListSubmitsRequest {
    #[prost(string, tag="1")]
    pub contest_slug: ::prost::alloc::string::String,
    /// 100 entries per 1 page
    #[prost(int32, tag="2")]
    pub page: i32,
    #[prost(enumeration="list_submits_request::SortBy", optional, tag="3")]
    pub sort_by: ::core::option::Option<i32>,
    #[prost(bool, optional, tag="4")]
    pub is_desc: ::core::option::Option<bool>,
    /// middle match
    #[prost(string, optional, tag="5")]
    pub group_name: ::core::option::Option<::prost::alloc::string::String>,
    #[prost(enumeration="Status", optional, tag="6")]
    pub status: ::core::option::Option<i32>,
    /// default: false
    #[prost(bool, optional, tag="7")]
    pub contains_guest: ::core::option::Option<bool>,
}
/// Nested message and enum types in `ListSubmitsRequest`.
pub mod list_submits_request {
    #[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
    #[repr(i32)]
    pub enum SortBy {
        SubmitedAt = 0,
        Score = 1,
    }
    impl SortBy {
        /// String value of the enum field names used in the ProtoBuf definition.
        ///
        /// The values are not transformed in any way and thus are considered stable
        /// (if the ProtoBuf definition does not change) and safe for programmatic use.
        pub fn as_str_name(&self) -> &'static str {
            match self {
                SortBy::SubmitedAt => "SUBMITED_AT",
                SortBy::Score => "SCORE",
            }
        }
        /// Creates an enum from field names used in the ProtoBuf definition.
        pub fn from_str_name(value: &str) -> ::core::option::Option<Self> {
            match value {
                "SUBMITED_AT" => Some(Self::SubmitedAt),
                "SCORE" => Some(Self::Score),
                _ => None,
            }
        }
    }
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct ListSubmitsResponse {
    /// NOTE: task_results will be empty
    #[prost(message, repeated, tag="1")]
    pub submits: ::prost::alloc::vec::Vec<Submit>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct CreateContestRequest {
    #[prost(string, tag="2")]
    pub title: ::prost::alloc::string::String,
    #[prost(message, optional, tag="4")]
    pub start_at: ::core::option::Option<::prost_types::Timestamp>,
    #[prost(message, optional, tag="5")]
    pub end_at: ::core::option::Option<::prost_types::Timestamp>,
    #[prost(int32, tag="6")]
    pub submit_limit: i32,
    #[prost(string, tag="8")]
    pub slug: ::prost::alloc::string::String,
    #[prost(enumeration="Validator", tag="11")]
    pub validator: i32,
    /// sec
    #[prost(int32, tag="12")]
    pub time_limit_per_task: i32,
    #[prost(oneof="create_contest_request::TagSelection", tags="9, 10")]
    pub tag_selection: ::core::option::Option<create_contest_request::TagSelection>,
}
/// Nested message and enum types in `CreateContestRequest`.
pub mod create_contest_request {
    #[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Oneof)]
    pub enum TagSelection {
        #[prost(message, tag="9")]
        Auto(super::TagSelectionLogicAuto),
        #[prost(message, tag="10")]
        Manual(super::TagSelectionLogicManual),
    }
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct CreateContestResponse {
    #[prost(message, optional, tag="1")]
    pub contest: ::core::option::Option<Contest>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct GetContestRequest {
    #[prost(string, tag="1")]
    pub contest_slug: ::prost::alloc::string::String,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct GetContestResponse {
    #[prost(message, optional, tag="1")]
    pub contest: ::core::option::Option<Contest>,
}
/// NOTE: cannot change slug and tag selection logic
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct UpdateContestRequest {
    #[prost(string, tag="1")]
    pub contest_slug: ::prost::alloc::string::String,
    #[prost(string, optional, tag="2")]
    pub title: ::core::option::Option<::prost::alloc::string::String>,
    #[prost(message, optional, tag="4")]
    pub start_at: ::core::option::Option<::prost_types::Timestamp>,
    #[prost(message, optional, tag="5")]
    pub end_at: ::core::option::Option<::prost_types::Timestamp>,
    #[prost(int32, optional, tag="6")]
    pub submit_limit: ::core::option::Option<i32>,
    #[prost(enumeration="Validator", optional, tag="9")]
    pub validator: ::core::option::Option<i32>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct UpdateContestResponse {
    #[prost(message, optional, tag="1")]
    pub contest: ::core::option::Option<Contest>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct GetRankingRequest {
    #[prost(string, tag="1")]
    pub contest_slug: ::prost::alloc::string::String,
    /// if it is true, return ranking which includes guests
    #[prost(bool, tag="2")]
    pub contain_guest: bool,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct GetRankingResponse {
    #[prost(message, repeated, tag="1")]
    pub records: ::prost::alloc::vec::Vec<get_ranking_response::Record>,
}
/// Nested message and enum types in `GetRankingResponse`.
pub mod get_ranking_response {
    #[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
    pub struct Record {
        #[prost(int32, tag="1")]
        pub rank: i32,
        #[prost(message, optional, tag="2")]
        pub group: ::core::option::Option<super::Group>,
        #[prost(int32, optional, tag="3")]
        pub score: ::core::option::Option<i32>,
    }
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct PingUnaryRequest {
    #[prost(string, tag="1")]
    pub ping: ::prost::alloc::string::String,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct PingUnaryResponse {
    #[prost(string, tag="1")]
    pub pong: ::prost::alloc::string::String,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct PingServerSideStreamingRequest {
    #[prost(string, tag="2")]
    pub ping: ::prost::alloc::string::String,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct PingServerSideStreamingResponse {
    #[prost(string, tag="2")]
    pub pong: ::prost::alloc::string::String,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct ListContestsRequest {
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct ListContestsResponse {
    #[prost(message, repeated, tag="2")]
    pub contests: ::prost::alloc::vec::Vec<Contest>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct VerifyTokenRequest {
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct VerifyTokenResponse {
    #[prost(bool, tag="1")]
    pub ok: bool,
    #[prost(string, tag="2")]
    pub message: ::prost::alloc::string::String,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct GetLatestSubmitRequest {
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct GetLatestSubmitResponse {
    #[prost(message, optional, tag="1")]
    pub submit: ::core::option::Option<Submit>,
}
include!("backend.tonic.rs");
// @@protoc_insertion_point(module)