package lib

import "github.com/robertkrimen/otto"

// ToValueFromString must parse value from string
func ToValueFromString(str string) otto.Value {
	v, _ := otto.ToValue(str)
	return v
}

// ToIntFromValue parses a value to integer.
// if an error occured, returns given default value
func ToIntFromValue(value otto.Value, defaultVal int) int {
	integer, err := value.ToInteger()
	if err != nil {
		return defaultVal
	}
	return int(integer)
}

// ToStringFromVM parse string from vm
func ToStringFromVM(vm *otto.Otto, key, defaultVal string) string {
	v, err := vm.Get(key)
	if err != nil {
		return defaultVal
	}
	return ToStringFromValue(v, defaultVal)
}

// ToStringFromValue parses a value to string
func ToStringFromValue(val otto.Value, defaultVal string) string {
	s, err := val.ToString()
	if err != nil {
		return defaultVal
	}
	return s
}

// GetEmptyObject returns {} (empty javascript object)
func GetEmptyObject(vm *otto.Otto) *otto.Object {
	object, _ := vm.Object("({})")
	return object
}
