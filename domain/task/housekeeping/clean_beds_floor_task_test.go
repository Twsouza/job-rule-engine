package task

import (
	"testing"

	"github.com/Twsouza/job-rule-engine/domain"
	"github.com/stretchr/testify/assert"
)

func TestCleanBedsFloor_AssertRule(t *testing.T) {
	cr := &CleanBedsFloor{}

	t.Run("should return true for valid job request", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: "Housekeeping",
			JobItem:    "Sheets",
			Locations:  []string{"Floor"},
		}

		result := cr.AssertRule(jobRequest)
		assert.True(t, result)
	})

	t.Run("should return false for invalid department", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: "Engineering",
			JobItem:    "Sheets",
			Locations:  []string{"Floor"},
		}

		result := cr.AssertRule(jobRequest)
		assert.False(t, result)
	})

	t.Run("should return false for invalid job item", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: "Housekeeping",
			JobItem:    "Pillow",
			Locations:  []string{"Floor"},
		}

		result := cr.AssertRule(jobRequest)
		assert.False(t, result)
	})

	t.Run("should return false for empty locations", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: "Housekeeping",
			JobItem:    "Sheets",
			Locations:  []string{},
		}

		result := cr.AssertRule(jobRequest)
		assert.False(t, result)
	})

	t.Run("should return false for missing floor location", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: "Housekeeping",
			JobItem:    "Sheets",
			Locations:  []string{"Room"},
		}

		result := cr.AssertRule(jobRequest)
		assert.False(t, result)
	})

	t.Run("should return false for missing job item", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: "Housekeeping",
			JobItem:    "",
			Locations:  []string{"Floor"},
		}

		result := cr.AssertRule(jobRequest)
		assert.False(t, result)
	})
}
