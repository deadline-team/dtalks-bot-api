// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package attachment

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

func easyjson76362c5bDecodeGithubComDeadlineTeamDtalksBotApiModelAttachment(in *jlexer.Lexer, out *Attachment) {
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
		case "author":
			if in.IsNull() {
				in.Skip()
				out.Author = nil
			} else {
				if out.Author == nil {
					out.Author = new(user.User)
				}
				easyjson76362c5bDecodeGithubComDeadlineTeamDtalksBotApiModelUser(in, out.Author)
			}
		case "fileName":
			out.FileName = string(in.String())
		case "mimeType":
			out.MimeType = string(in.String())
		case "size":
			out.Size = int64(in.Int64())
		case "path":
			out.Path = string(in.String())
		case "hash":
			out.Hash = string(in.String())
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
					var v1 interface{}
					if m, ok := v1.(easyjson.Unmarshaler); ok {
						m.UnmarshalEasyJSON(in)
					} else if m, ok := v1.(json.Unmarshaler); ok {
						_ = m.UnmarshalJSON(in.Raw())
					} else {
						v1 = in.Interface()
					}
					(out.Meta)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
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
func easyjson76362c5bEncodeGithubComDeadlineTeamDtalksBotApiModelAttachment(out *jwriter.Writer, in Attachment) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != "" {
		const prefix string = ",\"id\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.ID))
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
	if in.Author != nil {
		const prefix string = ",\"author\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson76362c5bEncodeGithubComDeadlineTeamDtalksBotApiModelUser(out, *in.Author)
	}
	if in.FileName != "" {
		const prefix string = ",\"fileName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.FileName))
	}
	if in.MimeType != "" {
		const prefix string = ",\"mimeType\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.MimeType))
	}
	if in.Size != 0 {
		const prefix string = ",\"size\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Size))
	}
	if in.Path != "" {
		const prefix string = ",\"path\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Path))
	}
	if in.Hash != "" {
		const prefix string = ",\"hash\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Hash))
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
			v2First := true
			for v2Name, v2Value := range in.Meta {
				if v2First {
					v2First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v2Name))
				out.RawByte(':')
				if m, ok := v2Value.(easyjson.Marshaler); ok {
					m.MarshalEasyJSON(out)
				} else if m, ok := v2Value.(json.Marshaler); ok {
					out.Raw(m.MarshalJSON())
				} else {
					out.Raw(json.Marshal(v2Value))
				}
			}
			out.RawByte('}')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Attachment) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson76362c5bEncodeGithubComDeadlineTeamDtalksBotApiModelAttachment(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Attachment) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson76362c5bEncodeGithubComDeadlineTeamDtalksBotApiModelAttachment(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Attachment) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson76362c5bDecodeGithubComDeadlineTeamDtalksBotApiModelAttachment(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Attachment) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson76362c5bDecodeGithubComDeadlineTeamDtalksBotApiModelAttachment(l, v)
}
func easyjson76362c5bDecodeGithubComDeadlineTeamDtalksBotApiModelUser(in *jlexer.Lexer, out *user.User) {
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
				easyjson76362c5bDecodeGithubComDeadlineTeamDtalksBotApiModelUser(in, out.Chief)
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
func easyjson76362c5bEncodeGithubComDeadlineTeamDtalksBotApiModelUser(out *jwriter.Writer, in user.User) {
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
		easyjson76362c5bEncodeGithubComDeadlineTeamDtalksBotApiModelUser(out, *in.Chief)
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
