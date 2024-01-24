package handler

import (
	"net/http"

	"github.com/Twsouza/job-rule-engine/application/dto"
	"github.com/Twsouza/job-rule-engine/domain/services"
	"github.com/gin-gonic/gin"
)

type JobRuleEngineHandler struct {
	JobService services.JobServiceInterface
}

func NewJobRuleEngineHandler(js services.JobServiceInterface) *JobRuleEngineHandler {
	return &JobRuleEngineHandler{
		JobService: js,
	}
}

func (jh *JobRuleEngineHandler) CreateJob(c *gin.Context) {
	req := &dto.JobRequestDto{}
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jobReq, errs := jh.JobService.LoadJob(req)
	if len(errs) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": errs})
		return
	}

	results := jh.JobService.CreateJob(jobReq)
	if len(results) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no rules matched for this job"})
		return
	}

	c.JSON(http.StatusCreated, results)
}
