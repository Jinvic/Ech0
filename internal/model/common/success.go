package model

import (
	i18n "github.com/lin-snow/ech0/internal/i18n"
)

// SUCCESS_MESSAGE 成功相关的消息常量
const (
	SUCCESS_MESSAGE = "success.success_message"
)

// Auth 成功相关常量
const (
	LOGIN_SUCCESS    = "success.login_success"
	REGISTER_SUCCESS = "success.register_success"
)

// Echo 成功相关常量
const (
	POST_ECHO_SUCCESS           = "success.post_echo_success"
	GET_ECHOS_BY_PAGE_SUCCESS   = "success.get_echos_by_page_success"
	DELETE_ECHO_SUCCESS         = "success.delete_echo_success"
	GET_TODAY_ECHOS_SUCCESS     = "success.get_today_echos_success"
	UPDATE_ECHO_SUCCESS         = "success.update_echo_success"
	LIKE_ECHO_SUCCESS           = "success.like_echo_success"
	GET_ECHO_BY_ID_SUCCESS      = "success.get_echo_by_id_success"
	GET_ALL_TAGS_SUCCESS        = "success.get_all_tags_success"
	DELETE_TAG_SUCCESS          = "success.delete_tag_success"
	GET_ECHOS_BY_TAG_ID_SUCCESS = "success.get_echos_by_tag_id_success"
)

// Common 成功相关常量
const (
	UPLOAD_SUCCESS             = "success.upload_success"
	DELETE_SUCCESS             = "success.delete_success"
	GET_STATUS_SUCCESS         = "success.get_status_success"
	GET_HEATMAP_SUCCESS        = "success.get_heatmap_success"
	GET_MUSIC_URL_SUCCESS      = "success.get_music_url_success"
	GET_HELLO_SUCCESS          = "success.get_hello_success"
	GET_S3_PRESIGN_URL_SUCCESS = "success.get_s3_presign_url_success"
	GET_METRICS_SUCCESS        = "success.get_metrics_success"
	GET_WEBSITE_TITLE_SUCCESS  = "success.get_website_title_success"
)

// Inbox 成功相关常量
const (
	GET_INBOX_LIST_SUCCESS   = "success.get_inbox_list_success"
	GET_UNREAD_INBOX_SUCCESS = "success.get_unread_inbox_success"
	MARK_INBOX_READ_SUCCESS  = "success.mark_inbox_read_success"
	DELETE_INBOX_SUCCESS     = "success.delete_inbox_success"
	CLEAR_INBOX_SUCCESS      = "success.clear_inbox_success"
)

// Setting 成功相关常量
const (
	GET_SETTINGS_SUCCESS              = "success.get_settings_success"
	UPDATE_SETTINGS_SUCCESS           = "success.update_settings_success"
	GET_COMMENT_SETTINGS_SUCCESS      = "success.get_comment_settings_success"
	UPDATE_COMMENT_SETTINGS_SUCCESS   = "success.update_comment_settings_success"
	GET_S3_SETTINGS_SUCCESS           = "success.get_s3_settings_success"
	UPDATE_S3_SETTINGS_SUCCESS        = "success.update_s3_settings_success"
	GET_OAUTH_SETTINGS_SUCCESS        = "success.get_oauth_settings_success"
	UPDATE_OAUTH_SETTINGS_SUCCESS     = "success.update_oauth_settings_success"
	GET_OAUTH2_STATUS_SUCCESS         = "success.get_oauth2_status_success"
	GET_WEBHOOK_SUCCESS               = "success.get_webhook_success"
	DELETE_WEBHOOK_SUCCESS            = "success.delete_webhook_success"
	UPDATE_WEBHOOK_SUCCESS            = "success.update_webhook_success"
	CREATE_WEBHOOK_SUCCESS            = "success.create_webhook_success"
	LIST_ACCESS_TOKENS_SUCCESS        = "success.list_access_tokens_success"
	CREATE_ACCESS_TOKEN_SUCCESS       = "success.create_access_token_success"
	DELETE_ACCESS_TOKEN_SUCCESS       = "success.delete_access_token_success"
	GET_FEDIVERSE_SETTINGS_SUCCESS    = "success.get_fediverse_settings_success"
	UPDATE_FEDIVERSE_SETTINGS_SUCCESS = "success.update_fediverse_settings_success"
	SCHEDULE_BACKUP_SUCCESS           = "success.schedule_backup_success"
)

// To do 成功相关常量
const (
	GET_TODO_LIST_SUCCESS = "success.get_todo_list_success"
	ADD_TODO_SUCCESS      = "success.add_todo_success"
	UPDATE_TODO_SUCCESS   = "success.update_todo_success"
	DELETE_TODO_SUCCESS   = "success.delete_todo_success"
)

// User 成功相关常量
const (
	UPDATE_USER_SUCCESS       = "success.update_user_success"
	GET_USER_SUCCESS          = "success.get_user_success"
	GET_USER_INFO_SUCCESS     = "success.get_user_info_success"
	DELETE_USER_SUCCESS       = "success.delete_user_success"
	BIND_GITHUB_SUCCESS       = "success.bind_github_success"
	GET_OAUTH_BINGURL_SUCCESS = "success.get_oauth_bingurl_success"
	GET_OAUTH_INFO_SUCCESS    = "success.get_oauth_info_success"
)

// Connect 成功相关常量
const (
	CONNECT_SUCCESS            = "success.connect_success"
	ADD_CONNECT_SUCCESS        = "success.add_connect_success"
	DELETE_CONNECT_SUCCESS     = "success.delete_connect_success"
	GET_CONNECT_INFO_SUCCESS   = "success.get_connect_info_success"
	GET_CONNECTED_LIST_SUCCESS = "success.get_connected_list_success"
)

// Backup 成功相关常量
const (
	BACKUP_SUCCESS        = "success.backup_success"
	EXPORT_BACKUP_SUCCESS = "success.export_backup_success"
	IMPORT_BACKUP_SUCCESS = "success.import_backup_success"
)

// Fediverse 成功相关常量
const (
	FEDIVERSE_SEARCH_ACTOR_SUCCESS      = "success.fediverse_search_actor_success"
	FEDIVERSE_FOLLOW_SUCCESS            = "success.fediverse_follow_success"
	FEDIVERSE_UNFOLLOW_SUCCESS          = "success.fediverse_unfollow_success"
	FEDIVERSE_LIKE_SUCCESS              = "success.fediverse_like_success"
	FEDIVERSE_UNDO_LIKE_SUCCESS         = "success.fediverse_undo_like_success"
	FEDIVERSE_GET_FOLLOW_STATUS_SUCCESS = "success.fediverse_get_follow_status_success"
	FEDIVERSE_GET_TIMELINE_SUCCESS      = "success.fediverse_get_timeline_success"
)

// Agent 成功相关常量
const (
	AGENT_GET_RECENT_SUCCESS = "success.agent_get_recent_success"
)

// GetSuccessMessage 获取本地化成功消息
func GetSuccessMessage(lang, messageID string) string {
	return i18n.T(lang, messageID)
}

// GetSuccessMessageWithData 获取本地化成功消息，包含模板数据
func GetSuccessMessageWithData(lang, messageID string, templateData map[string]any) string {
	return i18n.TWithData(lang, messageID, templateData)
}
