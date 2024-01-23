package factories

import (
	"github.com/Twsouza/job-rule-engine/domain/services"
	"github.com/Twsouza/job-rule-engine/domain/tasks"
	hk "github.com/Twsouza/job-rule-engine/domain/tasks/housekeeping"
	rs "github.com/Twsouza/job-rule-engine/domain/tasks/roomservice"
)

func NewJobService() *services.JobService {
	// TODO: create the JobApi implementation
	deliverJobItemLocation := rs.DeliverJobItemLocationTask{}
	deliverJobItemRoom := rs.DeliverJobItemRoomTask{}
	cleanBedsRoom := hk.CleanBedsRoom{}
	cleanBedsFloor := hk.CleanBedsFloor{}

	js := services.NewJobService([]tasks.JobTask{
		&deliverJobItemLocation,
		&deliverJobItemRoom,
		&cleanBedsRoom,
		&cleanBedsFloor,
	})

	return js
}
