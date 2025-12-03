package design_patterns

import (
	"context"
	"fmt"
	logger "log/slog"
)

// FilterInterface defines a filter with a Filter method
type FilterInterface[T any] interface {
	Filter(ctx context.Context, data *T) bool
	Name() string
}

// FilterBase provides a default implementation of FilterInterface
type FilterBase[T any] struct{}

func (b *FilterBase[T]) Filter(ctx context.Context, data *T) bool {
	return false
}

func (b *FilterBase[T]) Name() string {
	return "FilterBase"
}

// FilterHandler holds a list of filters to be filtered.
type FilterHandler[T any] struct {
	Filters []FilterInterface[T]
}

func (f *FilterHandler[T]) AddFilter(ctx context.Context, filter FilterInterface[T]) {
	f.Filters = append(f.Filters, filter)
}

func (f *FilterHandler[T]) Filter(ctx context.Context, data []*T) ([]*T, error) {
	if len(f.Filters) == 0 {
		return data, nil
	}

	passed := data

	for _, filter := range f.Filters {
		filtered := make([]*T, 0, len(passed))
		currentPassed := make([]*T, 0, len(passed))

		for _, item := range passed {
			if filter.Filter(ctx, item) {
				filtered = append(filtered, item)
			} else {
				currentPassed = append(currentPassed, item)
			}
		}

		logger.InfoContext(ctx, fmt.Sprintf("filter: %s, original: %d, passed: %d, failed: %d",
			filter.Name(),
			len(passed),
			len(currentPassed),
			len(filtered),
		))

		passed = currentPassed
	}

	return passed, nil
}
