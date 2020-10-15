package global

import (
	"github.com/kevinchfe/blog-service/pkg/logger"
	"github.com/kevinchfe/blog-service/pkg/setting"
)

var (
	ServerSetting *setting.ServerSettingS
	AppSetting *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS

	Logger *logger.Logger
)
