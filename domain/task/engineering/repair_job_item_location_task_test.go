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
			Department: &domain.Department{
				Name: "Engineering",
			},
			JobItem: &domain.JobItem{
				DisplayName: "TV",
			},
			Locations: []domain.Location{
				{
					LocationType: &domain.LocationType{
						DisplayName: "Room",
					},
				},
				{
					LocationType: &domain.LocationType{
						DisplayName: "Room",
					},
				},
			},
		}
		result := rj.AssertRule(jobRequest)
		assert.True(t, result)
	})

	t.Run("should return false for invalid department", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				Name: "Housekeeping",
			},
			JobItem: &domain.JobItem{
				DisplayName: "TV",
			},
			Locations: []domain.Location{
				{
					LocationType: &domain.LocationType{
						DisplayName: "Room",
					},
				},
			},
		}

		result := rj.AssertRule(jobRequest)
		assert.False(t, result)
	})

	t.Run("should return false for missing job item", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				Name: "Engineering",
			},
			JobItem: nil,
			Locations: []domain.Location{
				{
					LocationType: &domain.LocationType{
						DisplayName: "Room",
					},
				},
			},
		}

		result := rj.AssertRule(jobRequest)
		assert.False(t, result)
	})

	t.Run("should return false for empty locations", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				Name: "Engineering",
			},
			JobItem: &domain.JobItem{
				DisplayName: "TV",
			},
			Locations: []domain.Location{},
		}

		result := rj.AssertRule(jobRequest)
		assert.False(t, result)
	})
}
