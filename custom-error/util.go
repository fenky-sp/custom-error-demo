package customerror

import (
	"encoding/json"
	"reflect"
)

func convertContextualErrorDataToString(input any) string {
	if !allowDataAttachment {
		return ""
	}

	var output string

	data, containsInterface := processContextualErrorData(input)
	if containsInterface {
		output = "entire data is masked because it contains interface"
	} else {
		dataJsonBytes, _ := json.Marshal(data)
		output = string(dataJsonBytes)
	}

	return output
}

func processContextualErrorData(input any) (interface{}, bool) {
	var (
		output            interface{}
		containsInterface bool
	)

	if input == nil {
		return output, containsInterface
	}

	inputValue := reflect.ValueOf(input)

	inputBytes, err := json.Marshal(inputValue.Interface())
	if err != nil {
		return output, containsInterface
	}

	newValue := reflect.New(inputValue.Type()) // create zero value of same type as input value

	err = json.Unmarshal(inputBytes, newValue.Interface())
	if err != nil {
		return output, containsInterface
	}

	di := &dataIterator{}
	di.iterateData(newValue, "", "", processData)

	output = newValue.Interface()
	containsInterface = di.ContainsInterface

	return output, containsInterface
}

type errorData struct {
	errs []error
}

func (ed *errorData) getError(err error) {
	err = extractError(err)

	switch x := err.(type) {

	case interface{ Unwrap() error }:
		err = x.Unwrap()
		if err != nil {
			ed.getError(err) // check recursively if error is wrapped
		}

	case interface{ Unwrap() []error }:
		wrappedErrors := x.Unwrap()
		for _, wrappedError := range wrappedErrors {
			ed.getError(wrappedError)
		}

	default:
		ed.errs = append(ed.errs, err)

	}
}

func extractError(err error) error {
	switch err.(type) {
	case *metadata:
		if em, ok := err.(*metadata); ok {
			return em.err
		}
	}

	return err
}
