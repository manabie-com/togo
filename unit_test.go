package main

import (
	"context"
	beeCtx "github.com/beego/beego/v2/server/web/context"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
	"time"
	"togo/connections"
	"togo/helper"
	"togo/models"
)

func TestValidTask(t *testing.T) {
	assertions := assert.New(t)
	var tests = []struct {
		task models.Task
		expected bool
	}{
		{models.Task{
			Summary: "",
			Description: "",
			Assignee: "IYadf5AYZYZByyTTl1f5QqxOGx13",
			TaskDate: "2022-05-27",
		}, false},
		{models.Task{
			Summary: "Todo task",
			Description: "",
			Assignee: "",
			TaskDate: "2022-05-27",
		}, false},
		{models.Task{
			Summary: "Todo task",
			Description: "",
			Assignee: "IYadf5AYZYZByyTTl1f5QqxOGx13",
			TaskDate: "2022-05-32",
		}, false},
		{models.Task{
			Summary: "Todo task",
			Description: "",
			Assignee: "IYadf5AYZYZByyTTl1f5QqxOGx13",
			TaskDate: "2022-05-31",
		}, true},
	}
	for _, test := range tests {
		assertions.Equal(test.expected, test.task.IsValidTask())
	}
}

func TestConnection(t *testing.T) {
	assertions := assert.New(t)
	_, err := connections.Connect()
	assertions.Equal(nil, err)
}

func TestBeforeTaskHappy(t *testing.T) {
	assertions := assert.New(t)
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		log.Fatalf("%v", err)
	}
	ctx, cancel := context.WithTimeout(req.Context(), 10*time.Second)
	defer cancel()
	req = req.WithContext(ctx)
	bCtx := beeCtx.Context{
		Request: req,
	}
	accepted, _, err := helper.BeforeTask(&bCtx)
	assertions.Equal(nil, err)
	assertions.Equal(true, accepted)
}

func TestBeforeTaskDisconnected(t *testing.T) {
	assertions := assert.New(t)
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		log.Fatalf("%v", err)
	}
	ctx, cancel := context.WithTimeout(req.Context(), 1*time.Nanosecond)
	defer cancel()
	req = req.WithContext(ctx)
	bCtx := beeCtx.Context{
		Request: req,
	}
	accepted, _, err := helper.BeforeTask(&bCtx)
	assertions.NotEqual(nil, err)
	assertions.Equal(false, accepted)
}