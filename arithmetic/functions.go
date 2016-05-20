package arithmetic

func Add(x1,x2 interface{}) interface{}{
	// потом навренем, проверти типов и все такое
	return x1.(int)+x2.(int)
}

func Sub(x1,x2 interface{}) interface{}{
	// потом навренем, проверти типов и все такое
	return x1.(int)-x2.(int)
}

//func Add(a, b interface{}) (interface{}, error) {
//	value_a := reflect.ValueOf(a)
//	value_b := reflect.ValueOf(b)
//
//	if value_a.Kind() != value_b.Kind() {
//		return nil, fmt.Errorf("Different kinds, can't add them.")
//	}
//
//	switch value_a.Kind() {
//	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//		return value_a.Int() + value_b.Int(), nil
//	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
//		return value_a.Uint() + value_b.Uint(), nil
//	case reflect.Float32, reflect.Float64:
//		return value_a.Float() + value_b.Float(), nil
//	case reflect.String:
//		return value_a.String() + value_b.String(), nil
//	default:
//		return nil, fmt.Errorf("Type does not support addition.")
//	}
//}