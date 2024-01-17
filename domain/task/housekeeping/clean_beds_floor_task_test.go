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
			Department: &domain.Department{
				Name: "Housekeeping",
			},
			JobItem: &domain.JobItem{
				DisplayName: "Sheets",
			},
			Locations: []domain.Location{
				{
					LocationType: &domain.LocationType{
						DisplayName: "Floor",
					},
				},
			},
		}

		result := cr.AssertRule(jobRequest)
		assert.True(t, result)
	})

	t.Run("should return false for invalid department", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				Name: "Engineering",
			},
			JobItem: &domain.JobItem{
				DisplayName: "Sheets",
			},
			Locations: []domain.Location{
				{
					LocationType: &domain.LocationType{
						DisplayName: "Floor",
					},
				},
			},
		}

		result := cr.AssertRule(jobRequest)
		assert.False(t, result)
	})

	t.Run("should return false for invalid job item", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				Name: "Housekeeping",
			},
			JobItem: &domain.JobItem{
				DisplayName: "Pillow",
			},
			Locations: []domain.Location{
				{
					LocationType: &domain.LocationType{
						DisplayName: "Floor",
					},
				},
			},
		}

		result := cr.AssertRule(jobRequest)
		assert.False(t, result)
	})

	t.Run("should return false for empty locations", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				Name: "Housekeeping",
			},
			JobItem: &domain.JobItem{
				DisplayName: "Sheets",
			},
			Locations: []domain.Location{},
		}

		result := cr.AssertRule(jobRequest)
		assert.False(t, result)
	})

	t.Run("should return false for missing floor location", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				Name: "Housekeeping",
			},
			JobItem: &domain.JobItem{
				DisplayName: "Sheets",
			},
			Locations: []domain.Location{
				{
					LocationType: &domain.LocationType{
						DisplayName: "Room",
					},
				},
			},
		}

		result := cr.AssertRule(jobRequest)
		assert.False(t, result)
	})

	t.Run("should return false for missing job item", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				Name: "Housekeeping",
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

		result := cr.AssertRule(jobRequest)
		assert.False(t, result)
	})
}
