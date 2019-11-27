// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package config

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

func easyjsonCd93bc43Decode20192ShtobyShtoFileServiceConfig(in *jlexer.Lexer, out *Config) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "trello.service.port":
			out.Port = string(in.String())
		case "trello.service.storage.access.key":
			out.StorageAccessKey = string(in.String())
		case "trello.service.storage.secret.key":
			out.StorageSecretKey = string(in.String())
		case "trello.service.storage.region":
			out.StorageRegion = string(in.String())
		case "trello.service.storage.endpoint":
			out.StorageEndpoint = string(in.String())
		case "trello.service.storage.bucket":
			out.StorageBucket = string(in.String())
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
func easyjsonCd93bc43Encode20192ShtobyShtoFileServiceConfig(out *jwriter.Writer, in Config) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"trello.service.port\":"
		out.RawString(prefix[1:])
		out.String(string(in.Port))
	}
	{
		const prefix string = ",\"trello.service.storage.access.key\":"
		out.RawString(prefix)
		out.String(string(in.StorageAccessKey))
	}
	{
		const prefix string = ",\"trello.service.storage.secret.key\":"
		out.RawString(prefix)
		out.String(string(in.StorageSecretKey))
	}
	{
		const prefix string = ",\"trello.service.storage.region\":"
		out.RawString(prefix)
		out.String(string(in.StorageRegion))
	}
	{
		const prefix string = ",\"trello.service.storage.endpoint\":"
		out.RawString(prefix)
		out.String(string(in.StorageEndpoint))
	}
	{
		const prefix string = ",\"trello.service.storage.bucket\":"
		out.RawString(prefix)
		out.String(string(in.StorageBucket))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Config) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCd93bc43Encode20192ShtobyShtoFileServiceConfig(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Config) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCd93bc43Encode20192ShtobyShtoFileServiceConfig(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Config) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCd93bc43Decode20192ShtobyShtoFileServiceConfig(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Config) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCd93bc43Decode20192ShtobyShtoFileServiceConfig(l, v)
}