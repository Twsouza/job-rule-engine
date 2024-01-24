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

	deliverJobItemLocation := rs.DeliverJobItemLocationTask{
		API: optiSdk,
	}
	deliverJobItemRoom := rs.DeliverJobItemRoomTask{
		API: optiSdk,
	}
	cleanBedsRoom := hk.CleanBedsRoom{
		API: optiSdk,
	}
	cleanBedsFloor := hk.CleanBedsFloor{
		API: optiSdk,
	}
	taskList := []tasks.JobTask{
		&deliverJobItemLocation,
		&deliverJobItemRoom,
		&cleanBedsRoom,
		&cleanBedsFloor,
	}

	js := services.NewJobService(taskList, optiSdk)

	return js
}
