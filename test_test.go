package FieldProgramming

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

func TestTests(t *testing.T) {
	assert.Equal(t, gin.Default(), redis.Client{})
	assert.Equal(t, gorm.DB{}, postgres.Open(""))
}
