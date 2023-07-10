package error

import (
	"dootask-okr/i18n"

	"github.com/pkg/errors"
)

type WithError struct {
	Msg    string
	Detail interface{}
	Map    map[string]interface{}
	Err    error
}

func (e WithError) Error() string {
	content := ""
	if e.Detail != nil {
		content = i18n.GetErrMsg(e.Msg, map[string]any{"detail": e.Detail})
	} else if e.Map != nil {
		content = i18n.GetErrMsg(e.Msg, e.Map)
	} else {
		content = i18n.GetErrMsg(e.Msg, nil)
	}
	if content == "" {
		if e.Err != nil {
			return e.Err.Error()
		}
		return errors.New(e.Msg).Error()
	}
	return content
}

func New(Key string) WithError {
	return WithError{
		Msg:    Key,
		Detail: nil,
		Err:    nil,
	}
}

func WithDetail(Key string, detail any, err error) WithError {
	return WithError{
		Msg:    Key,
		Detail: detail,
		Err:    err,
	}
}

func WithErr(Key string, err error) WithError {
	return WithError{
		Msg:    Key,
		Detail: "",
		Err:    err,
	}
}

func WithMap(Key string, maps map[string]any, err error) WithError {
	return WithError{
		Msg: Key,
		Map: maps,
		Err: err,
	}
}
