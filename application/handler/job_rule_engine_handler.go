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
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.DepartmentID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "department_id is required"})
		return
	}
	if req.JobItemID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "job_item_id is required"})
		return
	}
	if len(req.LocationsID) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "locations_id is required"})
		return
	}

	jobReq, errs := jh.JobService.LoadJob(req)
	if len(errs) > 0 {
		errsStr := []string{}
		for _, err := range errs {
			errsStr = append(errsStr, err.Error())
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": errsStr})
		return
	}

	results := jh.JobService.CreateJob(jobReq)
	if len(results) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no rules matched for this job"})
		return
	}

	// Since a job request can match multiple rules, we return an array of job results.
	// Each job result contains the job request, the result of the rule, and any errors that occurred.
	// That's why we always return a 200 status code. To indicate that all rules were executed.
	// The consumer of this API can then decide what to do with the results.
	c.JSON(http.StatusOK, results)
}
