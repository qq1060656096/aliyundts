package parse

import "dhb168.com/x-store/aliyundts/dtsavro"

type DtsRecord dtsavro.Record

func (o *DtsRecord) GetBeforeImagesFields(f Formatter) (v []RecordField, err error) {
	beforeImagesArr := o.BeforeImages.ArrayUnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObject
	return o.GetImagesFieldsRaw(beforeImagesArr)
}

func (o *DtsRecord) GetAfterImagesFields() (v []RecordField, err error) {
	afterImagesArr := o.AfterImages.ArrayUnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObject
	return o.GetImagesFieldsRaw(afterImagesArr)
}

func (o *DtsRecord) GetImagesFieldsRaw(dtsImagesArray []*dtsavro.UnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObject) (v []RecordField, err error) {
	if dtsImagesArray == nil {
		return
	}

	l := len(o.Fields.ArrayField)
	v = make([]RecordField, 0, l)
	for k, field := range o.Fields.ArrayField {
		unionObject := NewFormatUnionObject(dtsImagesArray[k])
		fieldValue, rowErr := unionObject.GetValue()
		if rowErr != nil {
			err = rowErr
		}
		rv := RecordField {
			Name: field.Name,
			Value: fieldValue,
			Err: err,
		}
		v = append(v, rv)
	}
	return
}

func (o *DtsRecord) ImagesFieldsToMap(s []RecordField) (v map[string]interface{}, hasError bool, err error) {
	v = make(map[string]interface{})
	for _, field := range s {
		v[field.Name] = field.Value
		if !hasError && field.Err != nil {
			hasError = true
		}
	}
	return
}
