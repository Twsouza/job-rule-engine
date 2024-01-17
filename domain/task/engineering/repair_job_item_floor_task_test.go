package task

import (
	"testing"

	"github.com/Twsouza/job-rule-engine/domain"
	"github.com/stretchr/testify/assert"
)

func TestRepairJobItemFloor_AssertRule(t *testing.T) {
	rj := &RepairJobItemFloor{}

	t.Run("should return true for valid job request", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				Name: "Engineering",
			},
			JobItem: &domain.JobItem{
				DisplayName: "Air Conditioning",
			},
			Locations: []domain.Location{
				{
					LocationType: &domain.LocationType{
						DisplayName: "Floor",
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
				DisplayName: "Air Conditioning",
			},
			Locations: []domain.Location{
				{
					LocationType: &domain.LocationType{
						DisplayName: "Floor",
					},
				},
			},
		}

		result := rj.AssertRule(jobRequest)
		assert.False(t, result)
	})

	t.Run("should return false for empty job item", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				Name: "Engineering",
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

		result := rj.AssertRule(jobRequest)
		assert.False(t, result)
	})

	t.Run("should return false for invalid location", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				Name: "Engineering",
			},
			JobItem: &domain.JobItem{
				DisplayName: "Air Conditioning",
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

	t.Run("should return false for multiple locations", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				Name: "Engineering",
			},
			JobItem: &domain.JobItem{
				DisplayName: "Air Conditioning",
			},
			Locations: []domain.Location{
				{
					LocationType: &domain.LocationType{
						DisplayName: "Floor",
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
		assert.False(t, result)
	})
}
