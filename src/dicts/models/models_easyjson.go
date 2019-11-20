// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

import (
	customType "2019_2_Shtoby_shto/src/customType"
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

func easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels(in *jlexer.Lexer, out *User) {
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
		case "login":
			out.Login = string(in.String())
		case "password":
			out.Password = string(in.String())
		case "photo_id":
			if in.IsNull() {
				in.Skip()
				out.PhotoID = nil
			} else {
				if out.PhotoID == nil {
					out.PhotoID = new(customType.StringUUID)
				}
				*out.PhotoID = customType.StringUUID(in.String())
			}
		case "id":
			out.ID = customType.StringUUID(in.String())
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
func easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"login\":"
		out.RawString(prefix[1:])
		out.String(string(in.Login))
	}
	if in.Password != "" {
		const prefix string = ",\"password\":"
		out.RawString(prefix)
		out.String(string(in.Password))
	}
	if in.PhotoID != nil {
		const prefix string = ",\"photo_id\":"
		out.RawString(prefix)
		out.String(string(*in.PhotoID))
	}
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		out.String(string(in.ID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels(l, v)
}
func easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels1(in *jlexer.Lexer, out *Tag) {
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
		case "text":
			out.Text = string(in.String())
		case "id":
			out.ID = customType.StringUUID(in.String())
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
func easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels1(out *jwriter.Writer, in Tag) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"text\":"
		out.RawString(prefix[1:])
		out.String(string(in.Text))
	}
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		out.String(string(in.ID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Tag) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Tag) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Tag) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Tag) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels1(l, v)
}
func easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels2(in *jlexer.Lexer, out *Photo) {
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
		case "time_load":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.TimeLoad).UnmarshalJSON(data))
			}
		case "path":
			out.Path = string(in.String())
		case "id":
			out.ID = customType.StringUUID(in.String())
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
func easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels2(out *jwriter.Writer, in Photo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"time_load\":"
		out.RawString(prefix[1:])
		out.Raw((in.TimeLoad).MarshalJSON())
	}
	{
		const prefix string = ",\"path\":"
		out.RawString(prefix)
		out.String(string(in.Path))
	}
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		out.String(string(in.ID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Photo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Photo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Photo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Photo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels2(l, v)
}
func easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels3(in *jlexer.Lexer, out *ErrorResponse) {
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
		case "status":
			out.Status = int(in.Int())
		case "message":
			out.Message = string(in.String())
		case "error":
			out.Error = string(in.String())
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
func easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels3(out *jwriter.Writer, in ErrorResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Status))
	}
	{
		const prefix string = ",\"message\":"
		out.RawString(prefix)
		out.String(string(in.Message))
	}
	{
		const prefix string = ",\"error\":"
		out.RawString(prefix)
		out.String(string(in.Error))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ErrorResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ErrorResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ErrorResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ErrorResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels3(l, v)
}
func easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels4(in *jlexer.Lexer, out *Comment) {
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
		case "text":
			out.Text = string(in.String())
		case "card_id":
			out.CardID = string(in.String())
		case "user_id":
			out.UserID = string(in.String())
		case "id":
			out.ID = customType.StringUUID(in.String())
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
func easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels4(out *jwriter.Writer, in Comment) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"text\":"
		out.RawString(prefix[1:])
		out.String(string(in.Text))
	}
	{
		const prefix string = ",\"card_id\":"
		out.RawString(prefix)
		out.String(string(in.CardID))
	}
	{
		const prefix string = ",\"user_id\":"
		out.RawString(prefix)
		out.String(string(in.UserID))
	}
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		out.String(string(in.ID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Comment) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Comment) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Comment) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Comment) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels4(l, v)
}
func easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels5(in *jlexer.Lexer, out *CardsUserRequest) {
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
		case "users":
			if in.IsNull() {
				in.Skip()
				out.Users = nil
			} else {
				in.Delim('[')
				if out.Users == nil {
					if !in.IsDelim(']') {
						out.Users = make([]string, 0, 4)
					} else {
						out.Users = []string{}
					}
				} else {
					out.Users = (out.Users)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Users = append(out.Users, v1)
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
func easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels5(out *jwriter.Writer, in CardsUserRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"users\":"
		out.RawString(prefix[1:])
		if in.Users == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Users {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CardsUserRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CardsUserRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CardsUserRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CardsUserRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels5(l, v)
}
func easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels6(in *jlexer.Lexer, out *CardsUserAttachRequest) {
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
		case "user_id":
			out.UserID = customType.StringUUID(in.String())
		case "card_id":
			out.CardID = customType.StringUUID(in.String())
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
func easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels6(out *jwriter.Writer, in CardsUserAttachRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"user_id\":"
		out.RawString(prefix[1:])
		out.String(string(in.UserID))
	}
	{
		const prefix string = ",\"card_id\":"
		out.RawString(prefix)
		out.String(string(in.CardID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CardsUserAttachRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CardsUserAttachRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CardsUserAttachRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CardsUserAttachRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels6(l, v)
}
func easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels7(in *jlexer.Lexer, out *CardsBoardRequest) {
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
		case "boards":
			if in.IsNull() {
				in.Skip()
				out.Boards = nil
			} else {
				in.Delim('[')
				if out.Boards == nil {
					if !in.IsDelim(']') {
						out.Boards = make([]string, 0, 4)
					} else {
						out.Boards = []string{}
					}
				} else {
					out.Boards = (out.Boards)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.Boards = append(out.Boards, v4)
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
func easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels7(out *jwriter.Writer, in CardsBoardRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"boards\":"
		out.RawString(prefix[1:])
		if in.Boards == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Boards {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CardsBoardRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CardsBoardRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CardsBoardRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CardsBoardRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels7(l, v)
}
func easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels8(in *jlexer.Lexer, out *CardUsers) {
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
		case "card_id":
			out.CardID = customType.StringUUID(in.String())
		case "user_id":
			out.UserID = customType.StringUUID(in.String())
		case "id":
			out.ID = customType.StringUUID(in.String())
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
func easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels8(out *jwriter.Writer, in CardUsers) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"card_id\":"
		out.RawString(prefix[1:])
		out.String(string(in.CardID))
	}
	{
		const prefix string = ",\"user_id\":"
		out.RawString(prefix)
		out.String(string(in.UserID))
	}
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		out.String(string(in.ID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CardUsers) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CardUsers) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CardUsers) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CardUsers) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels8(l, v)
}
func easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels9(in *jlexer.Lexer, out *CardTagsRequest) {
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
		case "cards":
			if in.IsNull() {
				in.Skip()
				out.Cards = nil
			} else {
				in.Delim('[')
				if out.Cards == nil {
					if !in.IsDelim(']') {
						out.Cards = make([]string, 0, 4)
					} else {
						out.Cards = []string{}
					}
				} else {
					out.Cards = (out.Cards)[:0]
				}
				for !in.IsDelim(']') {
					var v7 string
					v7 = string(in.String())
					out.Cards = append(out.Cards, v7)
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
func easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels9(out *jwriter.Writer, in CardTagsRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"cards\":"
		out.RawString(prefix[1:])
		if in.Cards == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.Cards {
				if v8 > 0 {
					out.RawByte(',')
				}
				out.String(string(v9))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CardTagsRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CardTagsRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CardTagsRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CardTagsRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels9(l, v)
}
func easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels10(in *jlexer.Lexer, out *CardTags) {
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
		case "card_id":
			out.CardID = customType.StringUUID(in.String())
		case "tag_id":
			out.TagID = customType.StringUUID(in.String())
		case "id":
			out.ID = customType.StringUUID(in.String())
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
func easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels10(out *jwriter.Writer, in CardTags) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"card_id\":"
		out.RawString(prefix[1:])
		out.String(string(in.CardID))
	}
	{
		const prefix string = ",\"tag_id\":"
		out.RawString(prefix)
		out.String(string(in.TagID))
	}
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		out.String(string(in.ID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CardTags) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels10(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CardTags) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels10(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CardTags) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels10(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CardTags) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels10(l, v)
}
func easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels11(in *jlexer.Lexer, out *CardGroup) {
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
		case "name":
			out.Name = string(in.String())
		case "board_id":
			out.BoardID = customType.StringUUID(in.String())
		case "cards":
			if in.IsNull() {
				in.Skip()
				out.Cards = nil
			} else {
				in.Delim('[')
				if out.Cards == nil {
					if !in.IsDelim(']') {
						out.Cards = make([]Card, 0, 1)
					} else {
						out.Cards = []Card{}
					}
				} else {
					out.Cards = (out.Cards)[:0]
				}
				for !in.IsDelim(']') {
					var v10 Card
					(v10).UnmarshalEasyJSON(in)
					out.Cards = append(out.Cards, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "id":
			out.ID = customType.StringUUID(in.String())
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
func easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels11(out *jwriter.Writer, in CardGroup) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"board_id\":"
		out.RawString(prefix)
		out.String(string(in.BoardID))
	}
	{
		const prefix string = ",\"cards\":"
		out.RawString(prefix)
		if in.Cards == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v11, v12 := range in.Cards {
				if v11 > 0 {
					out.RawByte(',')
				}
				(v12).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		out.String(string(in.ID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CardGroup) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CardGroup) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CardGroup) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CardGroup) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels11(l, v)
}
func easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels12(in *jlexer.Lexer, out *Card) {
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
		case "caption":
			out.Caption = string(in.String())
		case "text":
			out.Text = string(in.String())
		case "priority":
			out.Priority = int(in.Int())
		case "card_group_id":
			out.CardGroupID = customType.StringUUID(in.String())
		case "comments":
			if in.IsNull() {
				in.Skip()
				out.Comments = nil
			} else {
				in.Delim('[')
				if out.Comments == nil {
					if !in.IsDelim(']') {
						out.Comments = make([]Comment, 0, 1)
					} else {
						out.Comments = []Comment{}
					}
				} else {
					out.Comments = (out.Comments)[:0]
				}
				for !in.IsDelim(']') {
					var v13 Comment
					(v13).UnmarshalEasyJSON(in)
					out.Comments = append(out.Comments, v13)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "tags":
			if in.IsNull() {
				in.Skip()
				out.Tags = nil
			} else {
				in.Delim('[')
				if out.Tags == nil {
					if !in.IsDelim(']') {
						out.Tags = make([]Tag, 0, 2)
					} else {
						out.Tags = []Tag{}
					}
				} else {
					out.Tags = (out.Tags)[:0]
				}
				for !in.IsDelim(']') {
					var v14 Tag
					(v14).UnmarshalEasyJSON(in)
					out.Tags = append(out.Tags, v14)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "users":
			if in.IsNull() {
				in.Skip()
				out.Users = nil
			} else {
				in.Delim('[')
				if out.Users == nil {
					if !in.IsDelim(']') {
						out.Users = make([]string, 0, 4)
					} else {
						out.Users = []string{}
					}
				} else {
					out.Users = (out.Users)[:0]
				}
				for !in.IsDelim(']') {
					var v15 string
					v15 = string(in.String())
					out.Users = append(out.Users, v15)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "deadline":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Deadline).UnmarshalJSON(data))
			}
		case "id":
			out.ID = customType.StringUUID(in.String())
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
func easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels12(out *jwriter.Writer, in Card) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"caption\":"
		out.RawString(prefix[1:])
		out.String(string(in.Caption))
	}
	{
		const prefix string = ",\"text\":"
		out.RawString(prefix)
		out.String(string(in.Text))
	}
	{
		const prefix string = ",\"priority\":"
		out.RawString(prefix)
		out.Int(int(in.Priority))
	}
	{
		const prefix string = ",\"card_group_id\":"
		out.RawString(prefix)
		out.String(string(in.CardGroupID))
	}
	{
		const prefix string = ",\"comments\":"
		out.RawString(prefix)
		if in.Comments == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v16, v17 := range in.Comments {
				if v16 > 0 {
					out.RawByte(',')
				}
				(v17).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"tags\":"
		out.RawString(prefix)
		if in.Tags == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v18, v19 := range in.Tags {
				if v18 > 0 {
					out.RawByte(',')
				}
				(v19).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"users\":"
		out.RawString(prefix)
		if in.Users == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v20, v21 := range in.Users {
				if v20 > 0 {
					out.RawByte(',')
				}
				out.String(string(v21))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"deadline\":"
		out.RawString(prefix)
		out.Raw((in.Deadline).MarshalJSON())
	}
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		out.String(string(in.ID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Card) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels12(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Card) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels12(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Card) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels12(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Card) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels12(l, v)
}
func easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels13(in *jlexer.Lexer, out *BoardsUserAttachRequest) {
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
		case "user_id":
			out.UserID = customType.StringUUID(in.String())
		case "board_id":
			out.BoardID = customType.StringUUID(in.String())
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
func easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels13(out *jwriter.Writer, in BoardsUserAttachRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"user_id\":"
		out.RawString(prefix[1:])
		out.String(string(in.UserID))
	}
	{
		const prefix string = ",\"board_id\":"
		out.RawString(prefix)
		out.String(string(in.BoardID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BoardsUserAttachRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels13(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BoardsUserAttachRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels13(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BoardsUserAttachRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels13(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BoardsUserAttachRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels13(l, v)
}
func easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels14(in *jlexer.Lexer, out *BoardUsers) {
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
		case "board_id":
			out.BoardID = customType.StringUUID(in.String())
		case "user_id":
			out.UserID = customType.StringUUID(in.String())
		case "id":
			out.ID = customType.StringUUID(in.String())
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
func easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels14(out *jwriter.Writer, in BoardUsers) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"board_id\":"
		out.RawString(prefix[1:])
		out.String(string(in.BoardID))
	}
	{
		const prefix string = ",\"user_id\":"
		out.RawString(prefix)
		out.String(string(in.UserID))
	}
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		out.String(string(in.ID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BoardUsers) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels14(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BoardUsers) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels14(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BoardUsers) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels14(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BoardUsers) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels14(l, v)
}
func easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels15(in *jlexer.Lexer, out *Board) {
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
		case "name":
			out.Name = string(in.String())
		case "users":
			if in.IsNull() {
				in.Skip()
				out.Users = nil
			} else {
				in.Delim('[')
				if out.Users == nil {
					if !in.IsDelim(']') {
						out.Users = make([]string, 0, 4)
					} else {
						out.Users = []string{}
					}
				} else {
					out.Users = (out.Users)[:0]
				}
				for !in.IsDelim(']') {
					var v22 string
					v22 = string(in.String())
					out.Users = append(out.Users, v22)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "card_groups":
			if in.IsNull() {
				in.Skip()
				out.CardGroups = nil
			} else {
				in.Delim('[')
				if out.CardGroups == nil {
					if !in.IsDelim(']') {
						out.CardGroups = make([]CardGroup, 0, 1)
					} else {
						out.CardGroups = []CardGroup{}
					}
				} else {
					out.CardGroups = (out.CardGroups)[:0]
				}
				for !in.IsDelim(']') {
					var v23 CardGroup
					(v23).UnmarshalEasyJSON(in)
					out.CardGroups = append(out.CardGroups, v23)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "id":
			out.ID = customType.StringUUID(in.String())
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
func easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels15(out *jwriter.Writer, in Board) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"users\":"
		out.RawString(prefix)
		if in.Users == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v24, v25 := range in.Users {
				if v24 > 0 {
					out.RawByte(',')
				}
				out.String(string(v25))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"card_groups\":"
		out.RawString(prefix)
		if in.CardGroups == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v26, v27 := range in.CardGroups {
				if v26 > 0 {
					out.RawByte(',')
				}
				(v27).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		out.String(string(in.ID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Board) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels15(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Board) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncode20192ShtobyShtoSrcDictsModels15(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Board) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels15(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Board) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecode20192ShtobyShtoSrcDictsModels15(l, v)
}
