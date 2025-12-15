package model

import i18n "github.com/lin-snow/ech0/internal/i18n"

// Panic Constants
const (
	INIT_LOGGER_PANIC          = "panic.init_logger_panic"
	READ_CONFIG_PANIC          = "panic.read_config_panic"
	CREATE_DB_PATH_PANIC       = "panic.create_db_path_panic"
	DATABASE_NOT_INITED        = "panic.database_not_inited"
	INIT_DATABASE_PANIC        = "panic.init_database_panic"
	MIGRATE_DB_PANIC           = "panic.migrate_db_panic"
	INIT_HANDLERS_PANIC        = "panic.init_handlers_panic"
	INIT_TASKER_PANIC          = "panic.init_tasker_panic"
	INIT_EVENT_REGISTRAR_PANIC = "panic.init_event_registrar_panic"
	GIN_RUN_FAILED             = "panic.gin_run_failed"
)

// GetPanicMessage 获取本地化panic消息
func GetPanicMessage(lang, messageID string) string {
	return i18n.T(lang, messageID)
}
