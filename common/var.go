package common

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var SugarLogger *zap.SugaredLogger
var DB *gorm.DB
var CTX = context.Background()
var RDB *redis.Client
