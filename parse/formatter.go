package parse

import (
	"github.com/qq1060656096/aliyundts/dtsavro"
)

type Formatter interface {
	Integer(r *dtsavro.Integer) (v interface{}, err error)
	Character(r *dtsavro.Character) (v interface{}, err error)
	Decimal(r *dtsavro.Decimal) (v interface{}, err error)
	Float(r *dtsavro.Float) (v interface{}, err error)
	Timestamp(r *dtsavro.Timestamp) (v interface{}, err error)
	DateTime(dt *dtsavro.DateTime) (v interface{}, err error)
	TimestampWithTimeZone(r *dtsavro.TimestampWithTimeZone) (v interface{}, err error)
	BinaryGeometry(r *dtsavro.BinaryGeometry) (v interface{}, err error)
	TextGeometry(r *dtsavro.TextGeometry) (v interface{}, err error)
	BinaryObject(r *dtsavro.BinaryObject) (v interface{}, err error)
	TextObject(r *dtsavro.TextObject) (v interface{}, err error)
	EmptyObject(r *dtsavro.EmptyObject) (v interface{}, err error)
}