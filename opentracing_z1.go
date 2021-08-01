package gormopentracing

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
)

const (
	parentSpanGormKey = "opentracingParentSpan"
)

// SetParentSpanToGorm sets parent span to gorm settings, returns cloned DB
func SetParentSpanToGorm(ctx context.Context, db *gorm.DB) *gorm.DB {
	if ctx == nil {
		return db
	}

	parentSpan := opentracing.SpanFromContext(ctx)
	if parentSpan == nil {
		return db
	}

	return db.Model("getInstance").InstanceSet(parentSpanGormKey, parentSpan)
}
