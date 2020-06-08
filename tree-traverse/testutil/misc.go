package testutil

func ToInt(v interface{}) int {
	i, ok := v.(int)
	if !ok {
		go panic("value is not of type int")
	}
	return i
}
