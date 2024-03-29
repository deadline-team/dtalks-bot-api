package service

/*
 * Copyright © 2023 - 2024, "DEADLINE TEAM" LLC
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
	"github.com/deadline-team/dtalks-bot-api/model"
	attachmentModel "github.com/deadline-team/dtalks-bot-api/model/attachment"
	"github.com/deadline-team/dtalks-bot-api/util"
	"io"
	"mime/multipart"
	"net/http"
	"time"
)

const attachmentBasePath = "/api/attachment/attachments"

var attachmentSrv AttachmentService

type AttachmentService interface {
	// GetAttachmentById
	// Метод для получения вложений по ID
	GetAttachmentById(ctx context.Context, attachmentId string) ([]byte, error)

	// GetAttachmentMetaById
	// Метод для получения меты вложений по ID
	GetAttachmentMetaById(ctx context.Context, attachmentId string) (*attachmentModel.Attachment, error)

	// CreateAttachment
	// Метод для создания вложений на сервере
	CreateAttachment(ctx context.Context, fileName string, data []byte) (*attachmentModel.Attachment, error)
}

type attachmentService struct {
	model.BotBaseParam
	httpClient *http.Client
}

func NewAttachmentService(botBaseParam model.BotBaseParam) AttachmentService {
	if attachmentSrv != nil {
		return attachmentSrv
	}
	attachmentSrv = &attachmentService{
		BotBaseParam: botBaseParam,
		httpClient:   &http.Client{Timeout: time.Second * 30},
	}
	return attachmentSrv
}

func (service *attachmentService) GetAttachmentById(ctx context.Context, attachmentId string) ([]byte, error) {
	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodGet, fmt.Sprintf("%s/%s", attachmentBasePath, attachmentId), nil)
	if err != nil {
		return nil, err
	}

	response, err := service.httpClient.Do(request)
	defer util.CloseChecker(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}

	return io.ReadAll(response.Body)
}

func (service *attachmentService) GetAttachmentMetaById(ctx context.Context, attachmentId string) (*attachmentModel.Attachment, error) {
	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodGet, fmt.Sprintf("%s/%s/meta", attachmentBasePath, attachmentId), nil)
	if err != nil {
		return nil, err
	}

	response, err := service.httpClient.Do(request)
	defer util.CloseChecker(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}

	var attachment attachmentModel.Attachment
	if err := json.NewDecoder(response.Body).Decode(&attachment); err != nil {
		return nil, err
	}

	return &attachment, err
}

func (service *attachmentService) CreateAttachment(ctx context.Context, fileName string, data []byte) (*attachmentModel.Attachment, error) {
	buf := new(bytes.Buffer)
	bw := multipart.NewWriter(buf)
	fw, err := bw.CreateFormFile("file", fileName)
	if err != nil {
		return nil, err
	}
	_, err = fw.Write(data)
	if err != nil {
		return nil, err
	}
	if err = bw.Close(); err != nil {
		return nil, err
	}

	request, err := util.CreateHttpRequest(ctx, service.BotBaseParam, http.MethodPost, attachmentBasePath, buf)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", bw.FormDataContentType())

	response, err := service.httpClient.Do(request)
	defer util.CloseChecker(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 201 {
		return nil, errors.New(response.Status)
	}

	var attachment attachmentModel.Attachment
	if err := json.NewDecoder(response.Body).Decode(&attachment); err != nil {
		return nil, err
	}

	return &attachment, err
}
