package dtalks_bot_api

/*
 * Copyright © 2023, "DEADLINE TEAM" LLC
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are not permitted.
 *
 * THIS SOFTWARE IS PROVIDED BY "DEADLINE TEAM" LLC "AS IS" AND ANY
 * EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
 * WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 * DISCLAIMED. IN NO EVENT SHALL "DEADLINE TEAM" LLC BE LIABLE FOR ANY
 * DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
 * (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
 * LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
 * ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 * (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
 * SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 *
 * No reproductions or distributions of this code is permitted without
 * written permission from "DEADLINE TEAM" LLC.
 * Do not reverse engineer or modify this code.
 *
 * © "DEADLINE TEAM" LLC, All rights reserved.
 */

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	taskModel "github.com/deadline-team/dtalks-bot-api/model/task"
	"net/http"
)

const taskBasePath = "/api/task/tasks"

func (client *botAPI) GetTaskById(ctx context.Context, taskId string, fields string) (*taskModel.Task, error) {
	request, err := client.createRequest(ctx, http.MethodGet, fmt.Sprintf("%s/%s", taskBasePath, taskId), nil)
	if err != nil {
		return nil, err
	}
	appendTaskQueryParams(request, taskModel.TaskFilter{}, fields)

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}

	var task *taskModel.Task
	if err := json.NewDecoder(response.Body).Decode(task); err != nil {
		return nil, err
	}
	if err = response.Body.Close(); err != nil {
		return nil, err
	}
	return task, nil
}

func (client *botAPI) GetTaskAll(ctx context.Context, filter taskModel.TaskFilter, fields string) ([]taskModel.Task, error) {
	request, err := client.createRequest(ctx, http.MethodGet, taskBasePath, nil)
	if err != nil {
		return nil, err
	}
	appendTaskQueryParams(request, filter, fields)

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}

	var tasks []taskModel.Task
	if err := json.NewDecoder(response.Body).Decode(&tasks); err != nil {
		return nil, err
	}
	if err = response.Body.Close(); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (client *botAPI) CreateTask(ctx context.Context, task taskModel.Task) (*taskModel.Task, error) {
	data, err := json.Marshal(&task)
	if err != nil {
		return nil, err
	}

	request, err := client.createRequest(ctx, http.MethodPost, taskBasePath, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 201 {
		return nil, errors.New(response.Status)
	}
	if err := json.NewDecoder(response.Body).Decode(&task); err != nil {
		return nil, err
	}
	if err = response.Body.Close(); err != nil {
		return nil, err
	}

	return &task, err
}

func (client *botAPI) UpdateTask(ctx context.Context, task taskModel.Task) (*taskModel.Task, error) {
	data, err := json.Marshal(&task)
	if err != nil {
		return nil, err
	}

	request, err := client.createRequest(ctx, http.MethodPut, taskBasePath, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}
	if err := json.NewDecoder(response.Body).Decode(&task); err != nil {
		return nil, err
	}
	if err = response.Body.Close(); err != nil {
		return nil, err
	}

	return &task, err
}

func (client *botAPI) DeleteTaskById(ctx context.Context, taskId string) error {
	request, err := client.createRequest(ctx, http.MethodDelete, fmt.Sprintf("%s/%s", taskBasePath, taskId), nil)
	if err != nil {
		return err
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		return errors.New(response.Status)
	}
	if err = response.Body.Close(); err != nil {
		return err
	}
	return nil
}

func (client *botAPI) ResolveTaskById(ctx context.Context, taskId string) error {
	request, err := client.createRequest(ctx, http.MethodPut, fmt.Sprintf("%s/%s/resolve", taskBasePath, taskId), nil)
	if err != nil {
		return err
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		return errors.New(response.Status)
	}
	if err = response.Body.Close(); err != nil {
		return err
	}
	return nil
}

func appendTaskQueryParams(request *http.Request, filter taskModel.TaskFilter, fields string) {
	if filter.IDs != nil && len(filter.IDs) > 0 {
		for _, id := range filter.IDs {
			request.Form.Add("ids", id)
		}
	}
	if filter.ConversationId != "" {
		request.Form.Set("conversationId", filter.ConversationId)
	}
	if filter.Resolved {
		request.Form.Set("resolved", fmt.Sprintf("%t", filter.Resolved))
	}
	if filter.Search != "" {
		request.Form.Set("search", filter.Search)
	}
	if fields != "" {
		request.Form.Set("fields", fields)
	}
}
