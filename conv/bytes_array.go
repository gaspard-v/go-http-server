package conv

type BytesArrayAdapter struct {
	value *[]byte
}

func CreateBytesArrayAdapter(value *[]byte) *BytesArrayAdapter {
	return &BytesArrayAdapter{value}
}

func (bytesArrayAdapter *BytesArrayAdapter) ToUint64() uint64 {
	int_size := 8
	array_size := len(*bytesArrayAdapter.value)
	var return_value uint64 = 0
	if int_size < array_size {
		return 0
	}
	for i := 0; i < array_size; i++ {
		value := (uint64)((*bytesArrayAdapter.value)[i])
		return_value += value << (8 * i)
	}
	return return_value
}
