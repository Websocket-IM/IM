package common

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var SugarLogger *zap.SugaredLogger
var DB *gorm.DB
