package user

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
	"github.com/deadline-team/dtalks-bot-api/model"
	"time"
)

type User struct {
	ID                string        `json:"id,omitempty"`
	Source            string        `json:"source,omitempty"`
	Username          string        `json:"username,omitempty"`
	FirstName         string        `json:"firstName,omitempty"`
	LastName          string        `json:"lastName,omitempty"`
	Email             string        `json:"email,omitempty"`
	Position          string        `json:"position,omitempty"`
	Avatar            *model.Avatar `json:"avatar,omitempty"`
	Birthday          *time.Time    `json:"birthday,omitempty"`
	PhoneNumber       int64         `json:"phoneNumber,omitempty"`
	City              string        `json:"city,omitempty"`
	Company           string        `json:"company,omitempty"`
	Department        string        `json:"department,omitempty"`
	Chief             *User         `json:"chief,omitempty"`
	LastActivity      *time.Time    `json:"lastActivity,omitempty"`
	Blocked           bool          `json:"blocked,omitempty"`
	TimeZone          int64         `json:"timeZone,omitempty"`
	CanChangePassword bool          `json:"canChangePassword,omitempty"`
	CanChangeAvatar   bool          `json:"canChangeAvatar,omitempty"`
}

type UserFilter struct {
	IDs       []string `form:"ids"`
	Username  string   `form:"username"`
	FirstName string   `form:"firstName"`
	LastName  string   `form:"lastName"`
	Email     string   `form:"email"`
	Search    string   `form:"search"`
}
