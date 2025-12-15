package model

import i18n "github.com/lin-snow/ech0/internal/i18n"

// ServerError 定义服务器错误信息
type ServerError struct {
	Msg string
	Err error
}

// 失败相关的常量
const (
	INVALID_FILE_PATH      = "error.invalid_file_path"
	INVALID_REQUEST_BODY   = "error.invalid_request_body"
	INVALID_PARAMS_BODY    = "error.invalid_params_body"
	INVALID_QUERY_PARAMS   = "error.invalid_query_params"
	INVALID_REQUEST_METHOD = "error.invalid_request_method"
)

// Auth 错误相关常量
const (
	USERNAME_OR_PASSWORD_NOT_BE_EMPTY = "error.username_or_password_not_be_empty"
	PASSWORD_INCORRECT                = "error.password_incorrect"
	USER_NOTFOUND                     = "error.user_notfound"
	USER_COUNT_EXCEED_LIMIT           = "error.user_count_exceed_limit"
	USERNAME_HAS_EXISTS               = "error.username_has_exists"
	TOKEN_NOT_FOUND                   = "error.token_not_found"
	TOKEN_NOT_VALID                   = "error.token_not_valid"
	TOKEN_PARSE_ERROR                 = "error.token_parse_error"
	USER_REGISTER_NOT_ALLOW           = "error.user_register_not_allow"
)

// Echo 错误相关常量
const (
	NO_PERMISSION_DENIED  = "error.no_permission_denied"
	ECHO_CAN_NOT_BE_EMPTY = "error.echo_can_not_be_empty"
	ECHO_NOT_FOUND        = "error.echo_not_found"
)

// Common 错误相关常量
const (
	NO_FILE_UPLOAD_ERROR   = "error.no_file_upload_error"
	NO_FILE_STORAGE_ERROR  = "error.no_file_storage_error"
	FILE_TYPE_NOT_ALLOWED  = "error.file_type_not_allowed"
	FILE_SIZE_EXCEED_LIMIT = "error.file_size_exceed_limit"
	IMAGE_NOT_FOUND        = "error.image_not_found"
	INVALID_PARAMS         = "error.invalid_params"
	SIGNUP_FIRST           = "error.signup_first"
	S3_NOT_ENABLED         = "error.s3_not_enabled"
	S3_NOT_CONFIGURED      = "error.s3_not_configured"
	S3_CONFIG_ERROR        = "error.s3_config_error"
)

// Inbox 错误相关常量
const (
	INBOX_NOT_FOUND = "error.inbox_not_found"
)

// User 错误相关常量
const (
	USERNAME_ALREADY_EXISTS        = "error.username_already_exists"
	FAILED_TO_GET_GITHUB_LOGIN_URL = "error.failed_to_get_github_login_url"
	FAILED_TO_GET_GOOGLE_LOGIN_URL = "error.failed_to_get_google_login_url"
	FAILED_TO_GET_QQ_LOGIN_URL     = "error.failed_to_get_qq_login_url"
	FAILED_TO_GET_CUSTOM_LOGIN_URL = "error.failed_to_get_custom_login_url"
	OAUTH2_NOT_CONFIGURED          = "error.oauth2_not_configured"
	OAUTH2_NOT_ENABLED             = "error.oauth2_not_enabled"
	NO_PERMISSION_BINDING_GITHUB   = "error.no_permission_binding_github"
	NO_PERMISSION_BINDING_GOOGLE   = "error.no_permission_binding_google"
	NO_PERMISSION_BINDING_QQ       = "error.no_permission_binding_qq"
	NO_PERMISSION_BINDING_CUSTOM   = "error.no_permission_binding_custom"
)

// TODO 错误相关常量
const (
	TODO_EXCEED_LIMIT = "error.todo_exceed_limit"
)

// Connect 错误相关常量
const (
	INVALID_CONNECTION_URL = "error.invalid_connection_url"
	CONNECT_HAS_EXISTS     = "error.connect_has_exists"
)

// Setting 错误相关常量
const (
	NO_SUCH_COMMENT_PROVIDER            = "error.no_such_comment_provider"
	WEBHOOK_NAME_OR_URL_CANNOT_BE_EMPTY = "error.webhook_name_or_url_cannot_be_empty"
	INVALID_CRON_EXPRESSION             = "error.invalid_cron_expression"
)

// Backup 错误相关常量
const (
	SNAPSHOT_UPLOAD_FAILED  = "error.snapshot_upload_failed"
	SNAPSHOT_RESTORE_FAILED = "error.snapshot_restore_failed"
	DATABASE_CLOSE_FAILED   = "error.database_close_failed"
)

// Fediverse 错误相关常量
const (
	GET_ACTOR_ERROR         = "error.get_actor_error"
	ACTIVEPUB_NOT_ENABLED   = "error.activepub_not_enabled"
	FEDIVERSE_INVALID_INPUT = "error.fediverse_invalid_input"
	FOLLOW_RELATION_MISSING = "error.follow_relation_missing"
)

// Agent 错误相关常量
const (
	AGENT_NOT_ENABLED        = "error.agent_not_enabled"
	AGENT_PROVIDER_NOT_FOUND = "error.agent_provider_not_found"
	AGENT_API_KEY_MISSING    = "error.agent_api_key_missing"
	AGENT_MODEL_MISSING      = "error.agent_model_missing"
	AGENT_SETTING_NOT_FOUND  = "error.agent_setting_not_found"
)

// GetErrorMessage 获取本地化错误消息
func GetErrorMessage(lang, messageID string) string {
	return i18n.T(lang, messageID)
}
