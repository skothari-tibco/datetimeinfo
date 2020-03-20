package datetimeinfo

import (
	"errors"
	"fmt"
	"os"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
	"github.com/project-flogo/core/support/log"
	"github.com/tkuchiki/parsetime"
)

const (
	WI_DATETIME_LOCATION         string = "FLOGO_DATETIME_LOCATION"
	WI_DATETIME_LOCATION_DEFAULT string = "UTC"
)

type DateTimeInfo struct {
}

func init() {
	function.Register(&DateTimeInfo{})
}

func (s *DateTimeInfo) Name() string {
	return "dateTimeInfo"
}

func (s *DateTimeInfo) GetCategory() string {
	return "dateTimeInfo"
}

func (s *DateTimeInfo) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString}, false
}

func (s *DateTimeInfo) Eval(params ...interface{}) (interface{}, error) {

	date, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("Format date first argument must be string")
	}
	p, err := parsetime.NewParseTime(GetLocation())
	if err != nil {
		log.RootLogger().Errorf("New time parser %s error %s", date, err.Error())
		return date, err
	}
	t, err := p.Parse(date)

	if err != nil {
		return nil, err
	}

	switch params[1].(string) {
	case "day":
		return t.Day(), nil
	case "hour":
		return t.Hour(), nil
	case "second":
		return t.Second(), nil
	case "minute":
		return t.Minute(), nil
	case "nanosecond":
		return t.Nanosecond(), nil
	default:
		return nil, errors.New("Operation not defineds")
	}

	return nil, nil
}

func GetLocation() string {
	location, ok := os.LookupEnv(WI_DATETIME_LOCATION)
	if ok && location != "" {
		return location
	}
	return WI_DATETIME_LOCATION_DEFAULT
}
