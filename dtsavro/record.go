// Code generated by github.com/actgardner/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCE:
 *     Record.avsc
 */
package dtsavro

import (
	"github.com/actgardner/gogen-avro/v7/compiler"
	"github.com/actgardner/gogen-avro/v7/vm"
	"github.com/actgardner/gogen-avro/v7/vm/types"
	"io"
)

type Record struct {
	// version infomation
	Version int32 `json:"version"`
	// unique id of this record in the whole stream
	Id int64 `json:"id"`
	// record log timestamp
	SourceTimestamp int64 `json:"sourceTimestamp"`
	// record source location information
	SourcePosition string `json:"sourcePosition"`
	// safe record source location information, use to recovery.
	SafeSourcePosition string `json:"safeSourcePosition"`
	// record transation id
	SourceTxid string `json:"sourceTxid"`
	// source dataource
	Source *Source `json:"source"`

	Operation Operation `json:"operation"`

	ObjectName *UnionNullString `json:"objectName"`
	// time when this record is processed along the stream dataflow
	ProcessTimestamps *UnionNullArrayLong `json:"processTimestamps"`
	// tags to identify properties of this record
	Tags map[string]string `json:"tags"`

	Fields *UnionNullStringArrayField `json:"fields"`

	BeforeImages *UnionNullStringArrayUnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObject `json:"beforeImages"`

	AfterImages *UnionNullStringArrayUnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObject `json:"afterImages"`
	// the timestamp in unit of millisecond that record is born in source
	BornTimestamp int64 `json:"bornTimestamp"`
}

const RecordAvroCRC64Fingerprint = "x\xc4\xc4yg\x16\xeb\xae"

func NewRecord() *Record {
	return &Record{}
}

func DeserializeRecord(r io.Reader) (*Record, error) {
	t := NewRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func DeserializeRecordFromSchema(r io.Reader, schema string) (*Record, error) {
	t := NewRecord()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func writeRecord(r *Record, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.Version, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.Id, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.SourceTimestamp, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.SourcePosition, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.SafeSourcePosition, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.SourceTxid, w)
	if err != nil {
		return err
	}
	err = writeSource(r.Source, w)
	if err != nil {
		return err
	}
	err = writeOperation(r.Operation, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ObjectName, w)
	if err != nil {
		return err
	}
	err = writeUnionNullArrayLong(r.ProcessTimestamps, w)
	if err != nil {
		return err
	}
	err = writeMapString(r.Tags, w)
	if err != nil {
		return err
	}
	err = writeUnionNullStringArrayField(r.Fields, w)
	if err != nil {
		return err
	}
	err = writeUnionNullStringArrayUnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObject(r.BeforeImages, w)
	if err != nil {
		return err
	}
	err = writeUnionNullStringArrayUnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObject(r.AfterImages, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.BornTimestamp, w)
	if err != nil {
		return err
	}
	return err
}

func (r *Record) Serialize(w io.Writer) error {
	return writeRecord(r, w)
}

func (r *Record) Schema() string {
	return "{\"fields\":[{\"doc\":\"version infomation\",\"name\":\"version\",\"type\":\"int\"},{\"doc\":\"unique id of this record in the whole stream\",\"name\":\"id\",\"type\":\"long\"},{\"doc\":\"record log timestamp\",\"name\":\"sourceTimestamp\",\"type\":\"long\"},{\"doc\":\"record source location information\",\"name\":\"sourcePosition\",\"type\":\"string\"},{\"default\":\"\",\"doc\":\"safe record source location information, use to recovery.\",\"name\":\"safeSourcePosition\",\"type\":\"string\"},{\"default\":\"\",\"doc\":\"record transation id\",\"name\":\"sourceTxid\",\"type\":\"string\"},{\"doc\":\"source dataource\",\"name\":\"source\",\"type\":{\"fields\":[{\"name\":\"sourceType\",\"type\":{\"name\":\"SourceType\",\"namespace\":\"com.alibaba.dts.formats.avro\",\"symbols\":[\"MySQL\",\"Oracle\",\"SQLServer\",\"PostgreSQL\",\"MongoDB\",\"Redis\",\"DB2\",\"PPAS\",\"DRDS\",\"HBASE\",\"HDFS\",\"FILE\",\"OTHER\"],\"type\":\"enum\"}},{\"doc\":\"source datasource version information\",\"name\":\"version\",\"type\":\"string\"}],\"name\":\"Source\",\"namespace\":\"com.alibaba.dts.formats.avro\",\"type\":\"record\"}},{\"name\":\"operation\",\"namespace\":\"com.alibaba.dts.formats.avro\",\"type\":{\"name\":\"Operation\",\"symbols\":[\"INSERT\",\"UPDATE\",\"DELETE\",\"DDL\",\"BEGIN\",\"COMMIT\",\"ROLLBACK\",\"ABORT\",\"HEARTBEAT\",\"CHECKPOINT\",\"COMMAND\",\"FILL\",\"FINISH\",\"CONTROL\",\"RDB\",\"NOOP\",\"INIT\"],\"type\":\"enum\"}},{\"default\":null,\"name\":\"objectName\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"doc\":\"time when this record is processed along the stream dataflow\",\"name\":\"processTimestamps\",\"type\":[\"null\",{\"items\":\"long\",\"type\":\"array\"}]},{\"default\":{},\"doc\":\"tags to identify properties of this record\",\"name\":\"tags\",\"type\":{\"type\":\"map\",\"values\":\"string\"}},{\"default\":null,\"name\":\"fields\",\"type\":[\"null\",\"string\",{\"items\":{\"fields\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataTypeNumber\",\"type\":\"int\"}],\"name\":\"Field\",\"namespace\":\"com.alibaba.dts.formats.avro\",\"type\":\"record\"},\"type\":\"array\"}]},{\"default\":null,\"name\":\"beforeImages\",\"type\":[\"null\",\"string\",{\"items\":[\"null\",{\"fields\":[{\"name\":\"precision\",\"type\":\"int\"},{\"name\":\"value\",\"type\":\"string\"}],\"name\":\"Integer\",\"namespace\":\"com.alibaba.dts.formats.avro\",\"type\":\"record\"},{\"fields\":[{\"name\":\"charset\",\"type\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"Character\",\"namespace\":\"com.alibaba.dts.formats.avro\",\"type\":\"record\"},{\"fields\":[{\"name\":\"value\",\"type\":\"string\"},{\"name\":\"precision\",\"type\":\"int\"},{\"name\":\"scale\",\"type\":\"int\"}],\"name\":\"Decimal\",\"namespace\":\"com.alibaba.dts.formats.avro\",\"type\":\"record\"},{\"fields\":[{\"name\":\"value\",\"type\":\"double\"},{\"name\":\"precision\",\"type\":\"int\"},{\"name\":\"scale\",\"type\":\"int\"}],\"name\":\"Float\",\"namespace\":\"com.alibaba.dts.formats.avro\",\"type\":\"record\"},{\"fields\":[{\"name\":\"timestamp\",\"type\":\"long\"},{\"name\":\"millis\",\"type\":\"int\"}],\"name\":\"Timestamp\",\"namespace\":\"com.alibaba.dts.formats.avro\",\"type\":\"record\"},{\"fields\":[{\"default\":null,\"name\":\"year\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"month\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"day\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"hour\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"minute\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"second\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"millis\",\"type\":[\"null\",\"int\"]}],\"name\":\"DateTime\",\"namespace\":\"com.alibaba.dts.formats.avro\",\"type\":\"record\"},{\"fields\":[{\"name\":\"value\",\"type\":\"com.alibaba.dts.formats.avro.DateTime\"},{\"name\":\"timezone\",\"type\":\"string\"}],\"name\":\"TimestampWithTimeZone\",\"namespace\":\"com.alibaba.dts.formats.avro\",\"type\":\"record\"},{\"fields\":[{\"name\":\"type\",\"type\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"BinaryGeometry\",\"namespace\":\"com.alibaba.dts.formats.avro\",\"type\":\"record\"},{\"fields\":[{\"name\":\"type\",\"type\":\"string\"},{\"name\":\"value\",\"type\":\"string\"}],\"name\":\"TextGeometry\",\"namespace\":\"com.alibaba.dts.formats.avro\",\"type\":\"record\"},{\"fields\":[{\"name\":\"type\",\"type\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"BinaryObject\",\"namespace\":\"com.alibaba.dts.formats.avro\",\"type\":\"record\"},{\"fields\":[{\"name\":\"type\",\"type\":\"string\"},{\"name\":\"value\",\"type\":\"string\"}],\"name\":\"TextObject\",\"namespace\":\"com.alibaba.dts.formats.avro\",\"type\":\"record\"},{\"name\":\"EmptyObject\",\"namespace\":\"com.alibaba.dts.formats.avro\",\"symbols\":[\"NULL\",\"NONE\"],\"type\":\"enum\"}],\"type\":\"array\"}]},{\"default\":null,\"name\":\"afterImages\",\"type\":[\"null\",\"string\",{\"items\":[\"null\",\"com.alibaba.dts.formats.avro.Integer\",\"com.alibaba.dts.formats.avro.Character\",\"com.alibaba.dts.formats.avro.Decimal\",\"com.alibaba.dts.formats.avro.Float\",\"com.alibaba.dts.formats.avro.Timestamp\",\"com.alibaba.dts.formats.avro.DateTime\",\"com.alibaba.dts.formats.avro.TimestampWithTimeZone\",\"com.alibaba.dts.formats.avro.BinaryGeometry\",\"com.alibaba.dts.formats.avro.TextGeometry\",\"com.alibaba.dts.formats.avro.BinaryObject\",\"com.alibaba.dts.formats.avro.TextObject\",\"com.alibaba.dts.formats.avro.EmptyObject\"],\"type\":\"array\"}]},{\"default\":0,\"doc\":\"the timestamp in unit of millisecond that record is born in source\",\"name\":\"bornTimestamp\",\"type\":\"long\"}],\"name\":\"com.alibaba.dts.formats.avro.Record\",\"type\":\"record\"}"
}

func (r *Record) SchemaName() string {
	return "com.alibaba.dts.formats.avro.Record"
}

func (_ *Record) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *Record) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *Record) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *Record) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *Record) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *Record) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *Record) SetString(v string)   { panic("Unsupported operation") }
func (_ *Record) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Record) Get(i int) types.Field {
	switch i {
	case 0:
		return &types.Int{Target: &r.Version}
	case 1:
		return &types.Long{Target: &r.Id}
	case 2:
		return &types.Long{Target: &r.SourceTimestamp}
	case 3:
		return &types.String{Target: &r.SourcePosition}
	case 4:
		return &types.String{Target: &r.SafeSourcePosition}
	case 5:
		return &types.String{Target: &r.SourceTxid}
	case 6:
		r.Source = NewSource()

		return r.Source
	case 7:
		return &OperationWrapper{Target: &r.Operation}
	case 8:
		r.ObjectName = NewUnionNullString()

		return r.ObjectName
	case 9:
		r.ProcessTimestamps = NewUnionNullArrayLong()

		return r.ProcessTimestamps
	case 10:
		r.Tags = make(map[string]string)

		return &MapStringWrapper{Target: &r.Tags}
	case 11:
		r.Fields = NewUnionNullStringArrayField()

		return r.Fields
	case 12:
		r.BeforeImages = NewUnionNullStringArrayUnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObject()

		return r.BeforeImages
	case 13:
		r.AfterImages = NewUnionNullStringArrayUnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObject()

		return r.AfterImages
	case 14:
		return &types.Long{Target: &r.BornTimestamp}
	}
	panic("Unknown field index")
}

func (r *Record) SetDefault(i int) {
	switch i {
	case 4:
		r.SafeSourcePosition = ""
		return
	case 5:
		r.SourceTxid = ""
		return
	case 8:
		r.ObjectName = nil
		return
	case 9:
		r.ProcessTimestamps = nil
		return
	case 10:

		return
	case 11:
		r.Fields = nil
		return
	case 12:
		r.BeforeImages = nil
		return
	case 13:
		r.AfterImages = nil
		return
	case 14:
		r.BornTimestamp = 0
		return
	}
	panic("Unknown field index")
}

func (r *Record) NullField(i int) {
	switch i {
	case 8:
		r.ObjectName = nil
		return
	case 9:
		r.ProcessTimestamps = nil
		return
	case 11:
		r.Fields = nil
		return
	case 12:
		r.BeforeImages = nil
		return
	case 13:
		r.AfterImages = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ *Record) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *Record) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *Record) Finalize()                        {}

func (_ *Record) AvroCRC64Fingerprint() []byte {
	return []byte(RecordAvroCRC64Fingerprint)
}