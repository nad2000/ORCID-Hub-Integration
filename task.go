package main

import (
	"log"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	ID          int      `json:"id,omitempty"`
	CompletedAt string   `json:"completed-at,omitempty"`
	CreatedAt   string   `json:"created-at,omitempty"`
	ExpiresAt   string   `json:"expires-at,omitempty"`
	Filename    string   `json:"filename,omitempty"`
	Status      string   `json:"status,omitempty"`
	Type        string   `json:"task-type,omitempty"`
	Records     []Record `json:"records"`
}

type Record struct {
	ID                  int    `json:"id,omitempty"`
	AffiliationType     string `json:"affiliation-type"`
	City                string `json:"city,omitempty"`
	Country             string `json:"country,omitempty"`
	Department          string `json:"department,omitempty"`
	DisambiguatedID     string `json:"disambiguated-id,omitempty"`
	DisambiguatedSource string `json:"disambiguated-source,omitempty"`
	Email               string `json:"email,omitempty"`
	EndDate             string `json:"end-date,omitempty"`
	ExternalID          string `json:"external-id,omitempty"`
	FirstName           string `json:"first-name,omitempty"`
	IsActive            bool   `json:"is-active,omitempty"`
	LastName            string `json:"last-name,omitempty"`
	Orcid               string `json:"orcid,omitempty"`
	Organisation        string `json:"organisation,omitempty"`
	ProcessedAt         string `json:"processed-at,omitempty"`
	PutCode             string `json:"put-code,omitempty"`
	Role                string `json:"role,omitempty"`
	StartDate           string `json:"start-date,omitempty"`
	State               string `json:"state,omitempty"`
	Status              string `json:"status,omitempty"`
}

func (t *Task) activateTask() {
	var task Task
	log.Printf("Activate the task %q (ID: %d)", t.Filename, t.ID)
	err := oh.Put("api/v1/tasks/"+strconv.Itoa(t.ID), map[string]string{"status": "ACTIVE"}, &task)
	if err != nil {
		panic(err)
	}
}

func newTask() {
	taskFilename := taskFilenamePrefix + strconv.FormatInt(time.Now().Unix(), 36) + ".json"
	var task = Task{Filename: taskFilename, Type: "AFFILIATION", Records: []Record{}}
	err := oh.Post("api/v1/affiliations?filename="+taskFilename, task, &task)
	if err != nil {
		panic(err)
	}
	taskID = task.ID
	taskCreatedAt, err = time.Parse("2006-01-02T15:04:05", task.CreatedAt)
	log.Printf("*** NEW TASK: %#v", task)
	taskSetUp <- true
}

// Either get the task ID or activate outstanding tasks and start a new one
func setupTask() {

	now := time.Now()
	if taskID == 0 {
		var tasks []Task
		// Make sure the access token acquired
		log.Println("=======================================================================================")
		<-gotAccessToken
		oh.Get("api/v1/tasks?type=AFFILIATION", &tasks)
		for _, t := range tasks {
			log.Printf("TASK: %#v", t)
			if t.Status == "ACTIVE" || t.CompletedAt != "" || !strings.HasPrefix(t.Filename, taskFilenamePrefix) {
				continue
			}
			createdAt, err := time.Parse("2006-01-02T15:04:05", t.CreatedAt)
			if err != nil {
				continue
			}
			if now.Sub(createdAt).Minutes() > 1 && len(t.Records) > 0 {
				go t.activateTask()
				continue
			}
			taskID = t.ID
			taskCreatedAt = createdAt
			taskRecordCount = len(t.Records)
			goto FOUND_TASK
		}
		go newTask()
		return

	} else if now.Sub(taskCreatedAt).Minutes() > 1 && taskRecordCount > 0 {
		var task = Task{ID: taskID}
		task.activateTask()
		go newTask()
		return
	}
FOUND_TASK:
	taskSetUp <- true
}
