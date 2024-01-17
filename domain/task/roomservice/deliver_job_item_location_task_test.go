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
			Department: &domain.Department{
				Name: "Room Service",
			},
			JobItem: &domain.JobItem{
				DisplayName: "Food",
			},
			Locations: []domain.Location{
				{
					LocationType: &domain.LocationType{
						DisplayName: "Floor",
					},
				},
			},
		}

		result := dj.AssertRule(jobRequest)
		assert.True(t, result)
	})

	t.Run("should return false for invalid department", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				Name: "Engineering",
			},
			JobItem: &domain.JobItem{
				DisplayName: "Food",
			},
			Locations: []domain.Location{
				{
					LocationType: &domain.LocationType{
						DisplayName: "Floor",
					},
				},
			},
		}

		result := dj.AssertRule(jobRequest)
		assert.False(t, result)
	})

	t.Run("should return false for missing job item", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				Name: "Room Service",
			},
			JobItem: nil,
			Locations: []domain.Location{
				{
					LocationType: &domain.LocationType{
						DisplayName: "Floor",
					},
				},
			},
		}

		result := dj.AssertRule(jobRequest)
		assert.False(t, result)
	})

	t.Run("should return false for empty locations", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				Name: "Room Service",
			},
			JobItem: &domain.JobItem{
				DisplayName: "Food",
			},
			Locations: []domain.Location{},
		}

		result := dj.AssertRule(jobRequest)
		assert.False(t, result)
	})
}
