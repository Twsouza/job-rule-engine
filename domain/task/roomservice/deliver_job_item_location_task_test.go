package task

import (
	"testing"

	"github.com/Twsouza/job-rule-engine/domain"
	"github.com/stretchr/testify/assert"
)

func TestDeliverJobItemLocation_AssertRule(t *testing.T) {
	dj := &DeliverJobItemLocationTask{}

	t.Run("should return true for valid job request", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: "Room Service",
			JobItem:    "Food",
			Locations:  []string{"Floor"},
		}

		result := dj.AssertRule(jobRequest)
		assert.True(t, result)
	})

	t.Run("should return false for invalid department", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: "Engineering",
			JobItem:    "Food",
			Locations:  []string{"Floor"},
		}

		result := dj.AssertRule(jobRequest)
		assert.False(t, result)
	})

	t.Run("should return false for missing job item", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: "Room Service",
			JobItem:    "",
			Locations:  []string{"Floor"},
		}

		result := dj.AssertRule(jobRequest)
		assert.False(t, result)
	})

	t.Run("should return false for empty locations", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: "Room Service",
			JobItem:    "Food",
			Locations:  []string{},
		}

		result := dj.AssertRule(jobRequest)
		assert.False(t, result)
	})
}
