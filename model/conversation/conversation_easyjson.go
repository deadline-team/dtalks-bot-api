// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package conversation

import (
	json "encoding/json"
	model "github.com/deadline-team/dtalks-bot-api/model"
	user "github.com/deadline-team/dtalks-bot-api/model/user"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	time "time"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonA5648bb1DecodeGithubComDeadlineTeamDtalksBotApiModelConversation(in *jlexer.Lexer, out *ConversationFilter) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "IDs":
			if in.IsNull() {
				in.Skip()
				out.IDs = nil
			} else {
				in.Delim('[')
				if out.IDs == nil {
					if !in.IsDelim(']') {
						out.IDs = make([]string, 0, 4)
					} else {
						out.IDs = []string{}
					}
				} else {
					out.IDs = (out.IDs)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.IDs = append(out.IDs, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "Name":
			out.Name = string(in.String())
		case "Visibility":
			out.Visibility = Visibility(in.String())
		case "Search":
			out.Search = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonA5648bb1EncodeGithubComDeadlineTeamDtalksBotApiModelConversation(out *jwriter.Writer, in ConversationFilter) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"IDs\":"
		out.RawString(prefix[1:])
		if in.IDs == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.IDs {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"Name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"Visibility\":"
		out.RawString(prefix)
		out.String(string(in.Visibility))
	}
	{
		const prefix string = ",\"Search\":"
		out.RawString(prefix)
		out.String(string(in.Search))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ConversationFilter) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA5648bb1EncodeGithubComDeadlineTeamDtalksBotApiModelConversation(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ConversationFilter) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA5648bb1EncodeGithubComDeadlineTeamDtalksBotApiModelConversation(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ConversationFilter) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA5648bb1DecodeGithubComDeadlineTeamDtalksBotApiModelConversation(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ConversationFilter) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA5648bb1DecodeGithubComDeadlineTeamDtalksBotApiModelConversation(l, v)
}
func easyjsonA5648bb1DecodeGithubComDeadlineTeamDtalksBotApiModelConversation1(in *jlexer.Lexer, out *Conversation) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = string(in.String())
		case "type":
			if in.IsNull() {
				in.Skip()
				out.Type = nil
			} else {
				in.Delim('[')
				if out.Type == nil {
					if !in.IsDelim(']') {
						out.Type = make([]ConversationDType, 0, 4)
					} else {
						out.Type = []ConversationDType{}
					}
				} else {
					out.Type = (out.Type)[:0]
				}
				for !in.IsDelim(']') {
					var v4 ConversationDType
					v4 = ConversationDType(in.String())
					out.Type = append(out.Type, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "name":
			out.Name = string(in.String())
		case "avatar":
			if in.IsNull() {
				in.Skip()
				out.Avatar = nil
			} else {
				if out.Avatar == nil {
					out.Avatar = new(model.Avatar)
				}
				(*out.Avatar).UnmarshalEasyJSON(in)
			}
		case "visibility":
			out.Visibility = Visibility(in.String())
		case "createDate":
			if in.IsNull() {
				in.Skip()
				out.CreateDate = nil
			} else {
				if out.CreateDate == nil {
					out.CreateDate = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.CreateDate).UnmarshalJSON(data))
				}
			}
		case "owner":
			if in.IsNull() {
				in.Skip()
				out.Owner = nil
			} else {
				if out.Owner == nil {
					out.Owner = new(user.User)
				}
				(*out.Owner).UnmarshalEasyJSON(in)
			}
		case "admins":
			if in.IsNull() {
				in.Skip()
				out.Admins = nil
			} else {
				in.Delim('[')
				if out.Admins == nil {
					if !in.IsDelim(']') {
						out.Admins = make([]*user.User, 0, 8)
					} else {
						out.Admins = []*user.User{}
					}
				} else {
					out.Admins = (out.Admins)[:0]
				}
				for !in.IsDelim(']') {
					var v5 *user.User
					if in.IsNull() {
						in.Skip()
						v5 = nil
					} else {
						if v5 == nil {
							v5 = new(user.User)
						}
						(*v5).UnmarshalEasyJSON(in)
					}
					out.Admins = append(out.Admins, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "members":
			if in.IsNull() {
				in.Skip()
				out.Members = nil
			} else {
				in.Delim('[')
				if out.Members == nil {
					if !in.IsDelim(']') {
						out.Members = make([]*user.User, 0, 8)
					} else {
						out.Members = []*user.User{}
					}
				} else {
					out.Members = (out.Members)[:0]
				}
				for !in.IsDelim(']') {
					var v6 *user.User
					if in.IsNull() {
						in.Skip()
						v6 = nil
					} else {
						if v6 == nil {
							v6 = new(user.User)
						}
						(*v6).UnmarshalEasyJSON(in)
					}
					out.Members = append(out.Members, v6)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "mutedMembers":
			if in.IsNull() {
				in.Skip()
				out.MutedMembers = nil
			} else {
				in.Delim('[')
				if out.MutedMembers == nil {
					if !in.IsDelim(']') {
						out.MutedMembers = make([]*user.User, 0, 8)
					} else {
						out.MutedMembers = []*user.User{}
					}
				} else {
					out.MutedMembers = (out.MutedMembers)[:0]
				}
				for !in.IsDelim(']') {
					var v7 *user.User
					if in.IsNull() {
						in.Skip()
						v7 = nil
					} else {
						if v7 == nil {
							v7 = new(user.User)
						}
						(*v7).UnmarshalEasyJSON(in)
					}
					out.MutedMembers = append(out.MutedMembers, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "messages":
			if in.IsNull() {
				in.Skip()
				out.Messages = nil
			} else {
				in.Delim('[')
				if out.Messages == nil {
					if !in.IsDelim(']') {
						out.Messages = make([]*Message, 0, 8)
					} else {
						out.Messages = []*Message{}
					}
				} else {
					out.Messages = (out.Messages)[:0]
				}
				for !in.IsDelim(']') {
					var v8 *Message
					if in.IsNull() {
						in.Skip()
						v8 = nil
					} else {
						if v8 == nil {
							v8 = new(Message)
						}
						(*v8).UnmarshalEasyJSON(in)
					}
					out.Messages = append(out.Messages, v8)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "pinned":
			if in.IsNull() {
				in.Skip()
				out.Pinned = nil
			} else {
				in.Delim('[')
				if out.Pinned == nil {
					if !in.IsDelim(']') {
						out.Pinned = make([]*Message, 0, 8)
					} else {
						out.Pinned = []*Message{}
					}
				} else {
					out.Pinned = (out.Pinned)[:0]
				}
				for !in.IsDelim(']') {
					var v9 *Message
					if in.IsNull() {
						in.Skip()
						v9 = nil
					} else {
						if v9 == nil {
							v9 = new(Message)
						}
						(*v9).UnmarshalEasyJSON(in)
					}
					out.Pinned = append(out.Pinned, v9)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "unreadCount":
			out.UnreadCount = int(in.Int())
		case "messageCount":
			out.MessageCount = int(in.Int())
		case "unreadThread":
			if in.IsNull() {
				in.Skip()
				out.UnreadThread = nil
			} else {
				in.Delim('[')
				if out.UnreadThread == nil {
					if !in.IsDelim(']') {
						out.UnreadThread = make([]*Message, 0, 8)
					} else {
						out.UnreadThread = []*Message{}
					}
				} else {
					out.UnreadThread = (out.UnreadThread)[:0]
				}
				for !in.IsDelim(']') {
					var v10 *Message
					if in.IsNull() {
						in.Skip()
						v10 = nil
					} else {
						if v10 == nil {
							v10 = new(Message)
						}
						(*v10).UnmarshalEasyJSON(in)
					}
					out.UnreadThread = append(out.UnreadThread, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonA5648bb1EncodeGithubComDeadlineTeamDtalksBotApiModelConversation1(out *jwriter.Writer, in Conversation) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != "" {
		const prefix string = ",\"id\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	if len(in.Type) != 0 {
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v11, v12 := range in.Type {
				if v11 > 0 {
					out.RawByte(',')
				}
				out.String(string(v12))
			}
			out.RawByte(']')
		}
	}
	if in.Name != "" {
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	if in.Avatar != nil {
		const prefix string = ",\"avatar\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Avatar).MarshalEasyJSON(out)
	}
	if in.Visibility != "" {
		const prefix string = ",\"visibility\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Visibility))
	}
	if in.CreateDate != nil {
		const prefix string = ",\"createDate\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((*in.CreateDate).MarshalJSON())
	}
	if in.Owner != nil {
		const prefix string = ",\"owner\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Owner).MarshalEasyJSON(out)
	}
	if len(in.Admins) != 0 {
		const prefix string = ",\"admins\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v13, v14 := range in.Admins {
				if v13 > 0 {
					out.RawByte(',')
				}
				if v14 == nil {
					out.RawString("null")
				} else {
					(*v14).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if len(in.Members) != 0 {
		const prefix string = ",\"members\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v15, v16 := range in.Members {
				if v15 > 0 {
					out.RawByte(',')
				}
				if v16 == nil {
					out.RawString("null")
				} else {
					(*v16).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if len(in.MutedMembers) != 0 {
		const prefix string = ",\"mutedMembers\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v17, v18 := range in.MutedMembers {
				if v17 > 0 {
					out.RawByte(',')
				}
				if v18 == nil {
					out.RawString("null")
				} else {
					(*v18).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if len(in.Messages) != 0 {
		const prefix string = ",\"messages\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v19, v20 := range in.Messages {
				if v19 > 0 {
					out.RawByte(',')
				}
				if v20 == nil {
					out.RawString("null")
				} else {
					(*v20).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if len(in.Pinned) != 0 {
		const prefix string = ",\"pinned\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v21, v22 := range in.Pinned {
				if v21 > 0 {
					out.RawByte(',')
				}
				if v22 == nil {
					out.RawString("null")
				} else {
					(*v22).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if in.UnreadCount != 0 {
		const prefix string = ",\"unreadCount\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.UnreadCount))
	}
	if in.MessageCount != 0 {
		const prefix string = ",\"messageCount\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.MessageCount))
	}
	if len(in.UnreadThread) != 0 {
		const prefix string = ",\"unreadThread\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v23, v24 := range in.UnreadThread {
				if v23 > 0 {
					out.RawByte(',')
				}
				if v24 == nil {
					out.RawString("null")
				} else {
					(*v24).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Conversation) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA5648bb1EncodeGithubComDeadlineTeamDtalksBotApiModelConversation1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Conversation) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA5648bb1EncodeGithubComDeadlineTeamDtalksBotApiModelConversation1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Conversation) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA5648bb1DecodeGithubComDeadlineTeamDtalksBotApiModelConversation1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Conversation) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA5648bb1DecodeGithubComDeadlineTeamDtalksBotApiModelConversation1(l, v)
}
