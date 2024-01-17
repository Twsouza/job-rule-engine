package task

import (
	"testing"

	"github.com/Twsouza/job-rule-engine/domain"
	"github.com/stretchr/testify/assert"
)

func TestRepairJobItemLocation_AssertRule(t *testing.T) {
	rj := &RepairJobItemLocation{}

	t.Run("should return true for valid job request", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: "Engineering",
			JobItem:    "TV",
			Locations:  []string{"Location1", "Location2"},
		}

		result := rj.AssertRule(jobRequest)
		assert.True(t, result)
	})

	t.Run("should return false for invalid department", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: "Housekeeping",
			JobItem:    "TV",
			Locations:  []string{"Location1", "Location2"},
		}

		result := rj.AssertRule(jobRequest)
		assert.False(t, result)
	})

	t.Run("should return false for missing job item", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: "Engineering",
			JobItem:    "",
			Locations:  []string{"Location1", "Location2"},
		}

		result := rj.AssertRule(jobRequest)
		assert.False(t, result)
	})

	t.Run("should return false for empty locations", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: "Engineering",
			JobItem:    "TV",
			Locations:  []string{},
		}

		result := rj.AssertRule(jobRequest)
		assert.False(t, result)
	})
}
