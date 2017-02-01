package filter

import (
	"github.com/TIBCOSoftware/flogo-lib/flow/activity"
	"github.com/op/go-logging"
	"fmt"
	"strconv"
)

const (
	input   	 = "input"
	datasource   = "datasource"
	datatype     = "datatype"
	minvalue 	 = "minvalue"
	maxvalue     = "maxvalue"
	interval     = "interval"
	intervaltype = "intervaltype"
)

var ifInput interface{}
var ifMinValue interface {}
var ifMaxValue interface {}
var valueTooLow = false
var valueTooHigh = false

// log is the default package logger
var log = logging.MustGetLogger("activity-tibco-rest")

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// init create & register activity
func init() {
	md := activity.NewMetadata(jsonMetadata)
	activity.Register(&MyActivity{metadata: md})
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error)  {

	//filter out by default
	context.SetOutput("usevalue", false)
	context.SetOutput("result", "")

	// check input value data type
	datatypeInput := context.GetInput(datatype)
	ivdatatype, ok := datatypeInput.(string)
	if !ok {
		context.SetOutput("reason", "DATATYPE_NOT_SET")
		return true, fmt.Errorf("Data type not set.")
	}

	// check the value matches the data type
	ifInput = validateValue(context, input, ivdatatype)
	if ifInput == nil {
		context.SetOutput("reason", "INPUT_INVALID")
		return true, fmt.Errorf("Invalid input data.")
	}
	
	// check if minimum value is set and apply filter
	cxMinValue := context.GetInput(minvalue)
	if cxMinValue != "" {
		// there is a minvalue assigned, now check for validity
		ifMinValue = validateValue(context, minvalue, ivdatatype)
				if ifMinValue == nil {
				context.SetOutput("reason", "MIN_VALUE_INVALID")
				return true, fmt.Errorf("Invalid minimum value.")
		} else {
			switch ivdatatype {
				case "int" : if ifInput.(int) < ifMinValue.(int) {valueTooLow = true}
				case "uint" : if ifInput.(uint) < ifMinValue.(uint) {valueTooLow = true}
				case "float32" : if ifInput.(float32) < ifMinValue.(float32) {valueTooLow = true}
			}
			if valueTooLow {
					context.SetOutput("reason", "VALUE_TOO_LOW")
					return true, nil
			}
		}		
	}

	// check if maximum value is set and apply filter
	cxMaxValue := context.GetInput(maxvalue)
	if cxMaxValue != "" {
		// there is a minvalue assigned, now check for validity
		ifMaxValue = validateValue(context, maxvalue, ivdatatype)
				if ifMinValue == nil {
				context.SetOutput("reason", "MAX_VALUE_INVALID")
				return true, fmt.Errorf("Invalid maximum value.")
		} else {
			switch ivdatatype {
				case "int" : if ifInput.(int) > ifMaxValue.(int) {valueTooHigh = true}
				case "uint" : if ifInput.(uint) > ifMaxValue.(uint) {valueTooHigh = true}
				case "float32" : if ifInput.(float32) > ifMaxValue.(float32) {valueTooHigh = true}
			}
			if valueTooHigh {
					context.SetOutput("reason", "VALUE_TOO_HIGH")
					return true, nil
			}
		}		
	}

	// When not filtered out, put the input data in output
	context.SetOutput("result", ifInput)
	context.SetOutput("usevalue", true)

	return true, nil
}

////////////////////////////////////////////////////////////////////////////////////////
// Utils
////////////////////////////////////////////////////////////////////////////////////////

func validateValue(context activity.Context, element string, datatype string) interface{}  {

	dataInput := context.GetInput(element)

	switch datatype {
		case "int": dataOutput, err := strconv.ParseInt(dataInput.(string), 10, strconv.IntSize)
			if err == nil {
				return int(dataOutput)
			}
		case "uint": dataOutput, err := strconv.ParseUint(dataInput.(string), 10, strconv.IntSize)
			if err == nil {
				return uint(dataOutput)
			}
		case "float32": dataOutput, err := strconv.ParseFloat(dataInput.(string), 32)
			if err == nil {
				return float32(dataOutput)
			}
		default:
			return nil
	}
	return nil
}