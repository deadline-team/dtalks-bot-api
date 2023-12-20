// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package conference

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson426b19f9DecodeGithubComDeadlineTeamDtalksBotApiModelConference(in *jlexer.Lexer, out *Room) {
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
			out.Id = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "conversationId":
			out.ConversationId = string(in.String())
		case "numParticipants":
			out.NumParticipants = uint32(in.Uint32())
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
func easyjson426b19f9EncodeGithubComDeadlineTeamDtalksBotApiModelConference(out *jwriter.Writer, in Room) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Id != "" {
		const prefix string = ",\"id\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Id))
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
	if in.ConversationId != "" {
		const prefix string = ",\"conversationId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ConversationId))
	}
	if in.NumParticipants != 0 {
		const prefix string = ",\"numParticipants\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint32(uint32(in.NumParticipants))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Room) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson426b19f9EncodeGithubComDeadlineTeamDtalksBotApiModelConference(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Room) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson426b19f9EncodeGithubComDeadlineTeamDtalksBotApiModelConference(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Room) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson426b19f9DecodeGithubComDeadlineTeamDtalksBotApiModelConference(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Room) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson426b19f9DecodeGithubComDeadlineTeamDtalksBotApiModelConference(l, v)
}