package factories

import (
	"os"

	"github.com/Twsouza/job-rule-engine/domain/services"
	"github.com/Twsouza/job-rule-engine/domain/tasks"
	hk "github.com/Twsouza/job-rule-engine/domain/tasks/housekeeping"
	rs "github.com/Twsouza/job-rule-engine/domain/tasks/roomservice"
	"github.com/Twsouza/job-rule-engine/infrastructure/sdk"
)

func NewJobService() *services.JobService {
	optiSdk, err := sdk.NewOptiiSdk(os.Getenv("OPTII_BASE_URL"), os.Getenv("OPTII_API_VERSION"), 3, nil)
	if err != nil {
		panic(err)
	}

	// TODO: create the JobApi implementation
	deliverJobItemLocation := rs.DeliverJobItemLocationTask{}
	deliverJobItemRoom := rs.DeliverJobItemRoomTask{}
	cleanBedsRoom := hk.CleanBedsRoom{}
	cleanBedsFloor := hk.CleanBedsFloor{}
	taskList := []tasks.JobTask{
		&deliverJobItemLocation,
		&deliverJobItemRoom,
		&cleanBedsRoom,
		&cleanBedsFloor,
	}

	js := services.NewJobService(taskList, optiSdk)

	return js
}
