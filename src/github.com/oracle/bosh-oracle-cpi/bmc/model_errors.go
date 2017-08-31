package bmc

import (
	"oracle/baremetal/core/models"
	"reflect"
)

// CoreModelErrorMsg extracts the failure msg  embedded
// inside the "Payload" member of a given error
func CoreModelErrorMsg(err error) string {
	embeddedMsg := ""
	concrete := reflect.Indirect(reflect.ValueOf(err))
	if concrete.Kind() == reflect.Struct {
		p := concrete.FieldByName("Payload")
		nil := reflect.Value{}
		if p != nil {
			p = reflect.Indirect(p)
			if merr, ok := p.Interface().(models.Error); ok {
				embeddedMsg = "[" + *merr.Message + "]"
			}
		}
	}
	return err.Error() + embeddedMsg
}
