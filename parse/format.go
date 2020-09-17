package parse

import (
	"github.com/qq1060656096/aliyundts/dtsavro"
	"fmt"
	"strconv"
	"time"
)

type Format struct {
	Formatter
}

func NewFormat() Formatter {
	return &Format{}
}

func (o *Format) Integer(r *dtsavro.Integer) (v interface{}, err error) {
	var nv int64
	if r == nil {
		v = &nv
		return
	}
	nv, err = strconv.ParseInt(r.Value, 10, 64)
	v = &nv
	return
}

func (o *Format) Character(r *dtsavro.Character) (v interface{}, err error) {
	var nv string
	if r == nil {
		v = &nv
		return
	}
	nv = string(r.Value)
	v = &nv
	return
}


func (o *Format) Decimal(r *dtsavro.Decimal) (v interface{}, err error) {
	var nv string
	if r == nil {
		v = &nv
		return
	}
	nv = string(r.Value)
	v = &nv
	return
}

func (o *Format) Float(r *dtsavro.Float) (v interface{}, err error) {
	var nv float64
	if r == nil {
		v = &nv
		return
	}
	nv = r.Value
	v = &nv
	return
}

func (o *Format) Timestamp(r *dtsavro.Timestamp) (v interface{}, err error) {
	var nv string
	if r == nil {
		v = &nv
		return
	}
	nv = time.Unix(r.Timestamp, 0).Format("2006-01-02 15:04:05")
	v = &nv
	return
}

func (o *Format) DateTime(dt *dtsavro.DateTime) (v interface{}, err error) {
	var s *string
	if dt == nil {
		return s, nil
	}
	years := make([]int32, 6)
	if dt.Year != nil {
		years[0] = dt.Year.Int
	} else {
		return
	}

	if dt.Month != nil {
		years[1] = dt.Month.Int
	}
	if dt.Day != nil {
		years[2] = dt.Day.Int
	}

	if dt.Hour != nil {
		years[3] = dt.Hour.Int
	}
	if dt.Minute != nil {
		years[4] = dt.Minute.Int
	}
	if dt.Second != nil {
		years[5] = dt.Second.Int
	}
	*s = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", years[0], years[1], years[2], years[3], years[4], years[5])
	v = s
	return
}

func (o *Format) TimestampWithTimeZone(r *dtsavro.TimestampWithTimeZone) (v interface{}, err error) {
	var nv string
	if r == nil {
		v = &nv
		return
	}
	return o.DateTime(r.Value)
}

func (o *Format) BinaryGeometry(r *dtsavro.BinaryGeometry) (v interface{}, err error) {
	var nv string
	if r == nil {
		v = &nv
		return
	}
	nv = string(r.Value)
	v = &nv
	return
}

func (o *Format) TextGeometry(r *dtsavro.TextGeometry) (v interface{}, err error) {
	var nv string
	if r == nil {
		v = &nv
		return
	}
	nv = string(r.Value)
	v = &nv
	return
}

func (o *Format) BinaryObject(r *dtsavro.BinaryObject) (v interface{}, err error) {
	var nv string
	if r == nil {
		v = &nv
		return
	}
	nv = string(r.Value)
	v = &nv
	return
}

func (o *Format) TextObject(r *dtsavro.TextObject) (v interface{}, err error) {
	var nv string
	if r == nil {
		v = &nv
		return
	}
	nv = r.Value
	v = &nv
	return
}

func (o *Format) EmptyObject(r *dtsavro.EmptyObject) (v interface{}, err error) {
	var nv string
	if r == nil {
		v = &nv
		return
	}
	nv = r.String()
	v = &nv
	return
}