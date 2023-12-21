// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package conversation

import (
	json "encoding/json"
	model "github.com/deadline-team/dtalks-bot-api/model"
	attachment "github.com/deadline-team/dtalks-bot-api/model/attachment"
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

func easyjson4086215fDecodeGithubComDeadlineTeamDtalksBotApiModelConversation(in *jlexer.Lexer, out *Message) {
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
		case "subType":
			out.SubType = MessageSubType(in.String())
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
		case "text":
			out.Text = string(in.String())
		case "author":
			if in.IsNull() {
				in.Skip()
				out.Author = nil
			} else {
				if out.Author == nil {
					out.Author = new(user.User)
				}
				easyjson4086215fDecodeGithubComDeadlineTeamDtalksBotApiModelUser(in, out.Author)
			}
		case "reply":
			if in.IsNull() {
				in.Skip()
				out.Reply = nil
			} else {
				if out.Reply == nil {
					out.Reply = new(Message)
				}
				(*out.Reply).UnmarshalEasyJSON(in)
			}
		case "forward":
			if in.IsNull() {
				in.Skip()
				out.Forward = nil
			} else {
				if out.Forward == nil {
					out.Forward = new(Message)
				}
				(*out.Forward).UnmarshalEasyJSON(in)
			}
		case "thread":
			if in.IsNull() {
				in.Skip()
				out.Thread = nil
			} else {
				in.Delim('[')
				if out.Thread == nil {
					if !in.IsDelim(']') {
						out.Thread = make([]*Message, 0, 8)
					} else {
						out.Thread = []*Message{}
					}
				} else {
					out.Thread = (out.Thread)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *Message
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(Message)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Thread = append(out.Thread, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "threadCount":
			out.ThreadCount = int(in.Int())
		case "threadUnreadCount":
			out.ThreadUnreadCount = int(in.Int())
		case "meta":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Meta = make(model.Meta)
				} else {
					out.Meta = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v2 interface{}
					if m, ok := v2.(easyjson.Unmarshaler); ok {
						m.UnmarshalEasyJSON(in)
					} else if m, ok := v2.(json.Unmarshaler); ok {
						_ = m.UnmarshalJSON(in.Raw())
					} else {
						v2 = in.Interface()
					}
					(out.Meta)[key] = v2
					in.WantComma()
				}
				in.Delim('}')
			}
		case "edited":
			out.Edited = bool(in.Bool())
		case "editDate":
			if in.IsNull() {
				in.Skip()
				out.EditDate = nil
			} else {
				if out.EditDate == nil {
					out.EditDate = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.EditDate).UnmarshalJSON(data))
				}
			}
		case "unread":
			if in.IsNull() {
				in.Skip()
				out.UnreadUsers = nil
			} else {
				in.Delim('[')
				if out.UnreadUsers == nil {
					if !in.IsDelim(']') {
						out.UnreadUsers = make([]*user.User, 0, 8)
					} else {
						out.UnreadUsers = []*user.User{}
					}
				} else {
					out.UnreadUsers = (out.UnreadUsers)[:0]
				}
				for !in.IsDelim(']') {
					var v3 *user.User
					if in.IsNull() {
						in.Skip()
						v3 = nil
					} else {
						if v3 == nil {
							v3 = new(user.User)
						}
						easyjson4086215fDecodeGithubComDeadlineTeamDtalksBotApiModelUser(in, v3)
					}
					out.UnreadUsers = append(out.UnreadUsers, v3)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "labels":
			if in.IsNull() {
				in.Skip()
				out.Labels = nil
			} else {
				in.Delim('[')
				if out.Labels == nil {
					if !in.IsDelim(']') {
						out.Labels = make([]*Label, 0, 8)
					} else {
						out.Labels = []*Label{}
					}
				} else {
					out.Labels = (out.Labels)[:0]
				}
				for !in.IsDelim(']') {
					var v4 *Label
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						if v4 == nil {
							v4 = new(Label)
						}
						(*v4).UnmarshalEasyJSON(in)
					}
					out.Labels = append(out.Labels, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "messageReactions":
			if in.IsNull() {
				in.Skip()
				out.MessageReactions = nil
			} else {
				in.Delim('[')
				if out.MessageReactions == nil {
					if !in.IsDelim(']') {
						out.MessageReactions = make([]*MessageReaction, 0, 8)
					} else {
						out.MessageReactions = []*MessageReaction{}
					}
				} else {
					out.MessageReactions = (out.MessageReactions)[:0]
				}
				for !in.IsDelim(']') {
					var v5 *MessageReaction
					if in.IsNull() {
						in.Skip()
						v5 = nil
					} else {
						if v5 == nil {
							v5 = new(MessageReaction)
						}
						easyjson4086215fDecodeGithubComDeadlineTeamDtalksBotApiModelConversation1(in, v5)
					}
					out.MessageReactions = append(out.MessageReactions, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "attachments":
			if in.IsNull() {
				in.Skip()
				out.Attachments = nil
			} else {
				in.Delim('[')
				if out.Attachments == nil {
					if !in.IsDelim(']') {
						out.Attachments = make([]*attachment.Attachment, 0, 8)
					} else {
						out.Attachments = []*attachment.Attachment{}
					}
				} else {
					out.Attachments = (out.Attachments)[:0]
				}
				for !in.IsDelim(']') {
					var v6 *attachment.Attachment
					if in.IsNull() {
						in.Skip()
						v6 = nil
					} else {
						if v6 == nil {
							v6 = new(attachment.Attachment)
						}
						(*v6).UnmarshalEasyJSON(in)
					}
					out.Attachments = append(out.Attachments, v6)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "read":
			out.Read = bool(in.Bool())
		case "readDate":
			if in.IsNull() {
				in.Skip()
				out.ReadDate = nil
			} else {
				if out.ReadDate == nil {
					out.ReadDate = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.ReadDate).UnmarshalJSON(data))
				}
			}
		case "callConferenceId":
			out.CallConferenceId = string(in.String())
		case "history":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.History = make(model.Meta)
				} else {
					out.History = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v7 interface{}
					if m, ok := v7.(easyjson.Unmarshaler); ok {
						m.UnmarshalEasyJSON(in)
					} else if m, ok := v7.(json.Unmarshaler); ok {
						_ = m.UnmarshalJSON(in.Raw())
					} else {
						v7 = in.Interface()
					}
					(out.History)[key] = v7
					in.WantComma()
				}
				in.Delim('}')
			}
		case "buttons":
			if in.IsNull() {
				in.Skip()
				out.Buttons = nil
			} else {
				in.Delim('[')
				if out.Buttons == nil {
					if !in.IsDelim(']') {
						out.Buttons = make(MessageButtons, 0, 1)
					} else {
						out.Buttons = MessageButtons{}
					}
				} else {
					out.Buttons = (out.Buttons)[:0]
				}
				for !in.IsDelim(']') {
					var v8 MessageButton
					easyjson4086215fDecodeGithubComDeadlineTeamDtalksBotApiModelConversation2(in, &v8)
					out.Buttons = append(out.Buttons, v8)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "deleted":
			out.Deleted = bool(in.Bool())
		case "deletedDate":
			if in.IsNull() {
				in.Skip()
				out.DeletedDate = nil
			} else {
				if out.DeletedDate == nil {
					out.DeletedDate = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.DeletedDate).UnmarshalJSON(data))
				}
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
func easyjson4086215fEncodeGithubComDeadlineTeamDtalksBotApiModelConversation(out *jwriter.Writer, in Message) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != "" {
		const prefix string = ",\"id\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	if in.SubType != "" {
		const prefix string = ",\"subType\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SubType))
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
	if in.Text != "" {
		const prefix string = ",\"text\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Text))
	}
	if in.Author != nil {
		const prefix string = ",\"author\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson4086215fEncodeGithubComDeadlineTeamDtalksBotApiModelUser(out, *in.Author)
	}
	if in.Reply != nil {
		const prefix string = ",\"reply\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Reply).MarshalEasyJSON(out)
	}
	if in.Forward != nil {
		const prefix string = ",\"forward\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Forward).MarshalEasyJSON(out)
	}
	if len(in.Thread) != 0 {
		const prefix string = ",\"thread\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v9, v10 := range in.Thread {
				if v9 > 0 {
					out.RawByte(',')
				}
				if v10 == nil {
					out.RawString("null")
				} else {
					(*v10).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if in.ThreadCount != 0 {
		const prefix string = ",\"threadCount\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.ThreadCount))
	}
	if in.ThreadUnreadCount != 0 {
		const prefix string = ",\"threadUnreadCount\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.ThreadUnreadCount))
	}
	if len(in.Meta) != 0 {
		const prefix string = ",\"meta\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v11First := true
			for v11Name, v11Value := range in.Meta {
				if v11First {
					v11First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v11Name))
				out.RawByte(':')
				if m, ok := v11Value.(easyjson.Marshaler); ok {
					m.MarshalEasyJSON(out)
				} else if m, ok := v11Value.(json.Marshaler); ok {
					out.Raw(m.MarshalJSON())
				} else {
					out.Raw(json.Marshal(v11Value))
				}
			}
			out.RawByte('}')
		}
	}
	if in.Edited {
		const prefix string = ",\"edited\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Edited))
	}
	if in.EditDate != nil {
		const prefix string = ",\"editDate\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((*in.EditDate).MarshalJSON())
	}
	if len(in.UnreadUsers) != 0 {
		const prefix string = ",\"unread\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v12, v13 := range in.UnreadUsers {
				if v12 > 0 {
					out.RawByte(',')
				}
				if v13 == nil {
					out.RawString("null")
				} else {
					easyjson4086215fEncodeGithubComDeadlineTeamDtalksBotApiModelUser(out, *v13)
				}
			}
			out.RawByte(']')
		}
	}
	if len(in.Labels) != 0 {
		const prefix string = ",\"labels\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v14, v15 := range in.Labels {
				if v14 > 0 {
					out.RawByte(',')
				}
				if v15 == nil {
					out.RawString("null")
				} else {
					(*v15).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if len(in.MessageReactions) != 0 {
		const prefix string = ",\"messageReactions\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v16, v17 := range in.MessageReactions {
				if v16 > 0 {
					out.RawByte(',')
				}
				if v17 == nil {
					out.RawString("null")
				} else {
					easyjson4086215fEncodeGithubComDeadlineTeamDtalksBotApiModelConversation1(out, *v17)
				}
			}
			out.RawByte(']')
		}
	}
	if len(in.Attachments) != 0 {
		const prefix string = ",\"attachments\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v18, v19 := range in.Attachments {
				if v18 > 0 {
					out.RawByte(',')
				}
				if v19 == nil {
					out.RawString("null")
				} else {
					(*v19).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if in.Read {
		const prefix string = ",\"read\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Read))
	}
	if in.ReadDate != nil {
		const prefix string = ",\"readDate\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((*in.ReadDate).MarshalJSON())
	}
	if in.CallConferenceId != "" {
		const prefix string = ",\"callConferenceId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.CallConferenceId))
	}
	if len(in.History) != 0 {
		const prefix string = ",\"history\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v20First := true
			for v20Name, v20Value := range in.History {
				if v20First {
					v20First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v20Name))
				out.RawByte(':')
				if m, ok := v20Value.(easyjson.Marshaler); ok {
					m.MarshalEasyJSON(out)
				} else if m, ok := v20Value.(json.Marshaler); ok {
					out.Raw(m.MarshalJSON())
				} else {
					out.Raw(json.Marshal(v20Value))
				}
			}
			out.RawByte('}')
		}
	}
	if len(in.Buttons) != 0 {
		const prefix string = ",\"buttons\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v21, v22 := range in.Buttons {
				if v21 > 0 {
					out.RawByte(',')
				}
				easyjson4086215fEncodeGithubComDeadlineTeamDtalksBotApiModelConversation2(out, v22)
			}
			out.RawByte(']')
		}
	}
	if in.Deleted {
		const prefix string = ",\"deleted\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Deleted))
	}
	if in.DeletedDate != nil {
		const prefix string = ",\"deletedDate\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((*in.DeletedDate).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Message) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4086215fEncodeGithubComDeadlineTeamDtalksBotApiModelConversation(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Message) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4086215fEncodeGithubComDeadlineTeamDtalksBotApiModelConversation(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Message) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4086215fDecodeGithubComDeadlineTeamDtalksBotApiModelConversation(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Message) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4086215fDecodeGithubComDeadlineTeamDtalksBotApiModelConversation(l, v)
}
func easyjson4086215fDecodeGithubComDeadlineTeamDtalksBotApiModelConversation2(in *jlexer.Lexer, out *MessageButton) {
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
		case "label":
			out.Label = string(in.String())
		case "data":
			out.Data = string(in.String())
		case "url":
			out.Url = string(in.String())
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
func easyjson4086215fEncodeGithubComDeadlineTeamDtalksBotApiModelConversation2(out *jwriter.Writer, in MessageButton) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Label != "" {
		const prefix string = ",\"label\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Label))
	}
	if in.Data != "" {
		const prefix string = ",\"data\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Data))
	}
	if in.Url != "" {
		const prefix string = ",\"url\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Url))
	}
	out.RawByte('}')
}
func easyjson4086215fDecodeGithubComDeadlineTeamDtalksBotApiModelConversation1(in *jlexer.Lexer, out *MessageReaction) {
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
		case "reaction":
			if in.IsNull() {
				in.Skip()
				out.Reaction = nil
			} else {
				if out.Reaction == nil {
					out.Reaction = new(Reaction)
				}
				easyjson4086215fDecodeGithubComDeadlineTeamDtalksBotApiModelConversation3(in, out.Reaction)
			}
		case "count":
			out.Count = int(in.Int())
		case "users":
			if in.IsNull() {
				in.Skip()
				out.Users = nil
			} else {
				in.Delim('[')
				if out.Users == nil {
					if !in.IsDelim(']') {
						out.Users = make([]*user.User, 0, 8)
					} else {
						out.Users = []*user.User{}
					}
				} else {
					out.Users = (out.Users)[:0]
				}
				for !in.IsDelim(']') {
					var v23 *user.User
					if in.IsNull() {
						in.Skip()
						v23 = nil
					} else {
						if v23 == nil {
							v23 = new(user.User)
						}
						easyjson4086215fDecodeGithubComDeadlineTeamDtalksBotApiModelUser(in, v23)
					}
					out.Users = append(out.Users, v23)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "reactedByMe":
			out.ReactedByMe = bool(in.Bool())
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
func easyjson4086215fEncodeGithubComDeadlineTeamDtalksBotApiModelConversation1(out *jwriter.Writer, in MessageReaction) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != "" {
		const prefix string = ",\"id\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	if in.Reaction != nil {
		const prefix string = ",\"reaction\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson4086215fEncodeGithubComDeadlineTeamDtalksBotApiModelConversation3(out, *in.Reaction)
	}
	if in.Count != 0 {
		const prefix string = ",\"count\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Count))
	}
	if len(in.Users) != 0 {
		const prefix string = ",\"users\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v24, v25 := range in.Users {
				if v24 > 0 {
					out.RawByte(',')
				}
				if v25 == nil {
					out.RawString("null")
				} else {
					easyjson4086215fEncodeGithubComDeadlineTeamDtalksBotApiModelUser(out, *v25)
				}
			}
			out.RawByte(']')
		}
	}
	if in.ReactedByMe {
		const prefix string = ",\"reactedByMe\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.ReactedByMe))
	}
	out.RawByte('}')
}
func easyjson4086215fDecodeGithubComDeadlineTeamDtalksBotApiModelConversation3(in *jlexer.Lexer, out *Reaction) {
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
		case "value":
			out.Value = string(in.String())
		case "invisible":
			out.Invisible = bool(in.Bool())
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
func easyjson4086215fEncodeGithubComDeadlineTeamDtalksBotApiModelConversation3(out *jwriter.Writer, in Reaction) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != "" {
		const prefix string = ",\"id\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	if in.Value != "" {
		const prefix string = ",\"value\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Value))
	}
	if in.Invisible {
		const prefix string = ",\"invisible\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Invisible))
	}
	out.RawByte('}')
}
func easyjson4086215fDecodeGithubComDeadlineTeamDtalksBotApiModelUser(in *jlexer.Lexer, out *user.User) {
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
		case "source":
			out.Source = string(in.String())
		case "username":
			out.Username = string(in.String())
		case "firstName":
			out.FirstName = string(in.String())
		case "lastName":
			out.LastName = string(in.String())
		case "email":
			out.Email = string(in.String())
		case "position":
			out.Position = string(in.String())
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
		case "birthday":
			if in.IsNull() {
				in.Skip()
				out.Birthday = nil
			} else {
				if out.Birthday == nil {
					out.Birthday = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.Birthday).UnmarshalJSON(data))
				}
			}
		case "phoneNumber":
			out.PhoneNumber = int64(in.Int64())
		case "city":
			out.City = string(in.String())
		case "company":
			out.Company = string(in.String())
		case "department":
			out.Department = string(in.String())
		case "chief":
			if in.IsNull() {
				in.Skip()
				out.Chief = nil
			} else {
				if out.Chief == nil {
					out.Chief = new(user.User)
				}
				easyjson4086215fDecodeGithubComDeadlineTeamDtalksBotApiModelUser(in, out.Chief)
			}
		case "lastActivity":
			if in.IsNull() {
				in.Skip()
				out.LastActivity = nil
			} else {
				if out.LastActivity == nil {
					out.LastActivity = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.LastActivity).UnmarshalJSON(data))
				}
			}
		case "blocked":
			out.Blocked = bool(in.Bool())
		case "timeZone":
			out.TimeZone = int64(in.Int64())
		case "canChangePassword":
			out.CanChangePassword = bool(in.Bool())
		case "canChangeAvatar":
			out.CanChangeAvatar = bool(in.Bool())
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
func easyjson4086215fEncodeGithubComDeadlineTeamDtalksBotApiModelUser(out *jwriter.Writer, in user.User) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != "" {
		const prefix string = ",\"id\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	if in.Source != "" {
		const prefix string = ",\"source\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Source))
	}
	if in.Username != "" {
		const prefix string = ",\"username\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Username))
	}
	if in.FirstName != "" {
		const prefix string = ",\"firstName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.FirstName))
	}
	if in.LastName != "" {
		const prefix string = ",\"lastName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.LastName))
	}
	if in.Email != "" {
		const prefix string = ",\"email\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Email))
	}
	if in.Position != "" {
		const prefix string = ",\"position\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Position))
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
	if in.Birthday != nil {
		const prefix string = ",\"birthday\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((*in.Birthday).MarshalJSON())
	}
	if in.PhoneNumber != 0 {
		const prefix string = ",\"phoneNumber\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.PhoneNumber))
	}
	if in.City != "" {
		const prefix string = ",\"city\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.City))
	}
	if in.Company != "" {
		const prefix string = ",\"company\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Company))
	}
	if in.Department != "" {
		const prefix string = ",\"department\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Department))
	}
	if in.Chief != nil {
		const prefix string = ",\"chief\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson4086215fEncodeGithubComDeadlineTeamDtalksBotApiModelUser(out, *in.Chief)
	}
	if in.LastActivity != nil {
		const prefix string = ",\"lastActivity\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((*in.LastActivity).MarshalJSON())
	}
	if in.Blocked {
		const prefix string = ",\"blocked\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Blocked))
	}
	if in.TimeZone != 0 {
		const prefix string = ",\"timeZone\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.TimeZone))
	}
	if in.CanChangePassword {
		const prefix string = ",\"canChangePassword\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.CanChangePassword))
	}
	if in.CanChangeAvatar {
		const prefix string = ",\"canChangeAvatar\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.CanChangeAvatar))
	}
	out.RawByte('}')
}
