package eagofs

import "os"

// CreateFile creates a file given name
func CreateFile(name, content string) error {
	file, err := os.Create(name)
	if err != nil {
		return err
		//panic(eagrors.NewErrorWithCause(err, "can not create eago.json"))
	}
	_, err = file.Write([]byte(content))
	if err != nil {
		return err
		//panic(eagrors.NewErrorWithCause(err, "can not create eago.json"))
	}
	return nil
}
