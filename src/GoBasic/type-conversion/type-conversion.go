package typeConversion

import stringToInt "GoBasic/type-conversion/string-to-int"
import stringToFloat "GoBasic/type-conversion/string-to-float"
import stringToBool "GoBasic/type-conversion/string-to-bool"
import boolToString "GoBasic/type-conversion/bool-to-string"
import floatToString "GoBasic/type-conversion/float-to-string"
import intToString "GoBasic/type-conversion/int-to-string"
import intToSpecificInt "GoBasic/type-conversion/int-to-specific-int"
import floatToSpecificFloat "GoBasic/type-conversion/float-to-specific-float"
import intToFloat "GoBasic/type-conversion/int-to-float"

func TypeConversion() {
	// Type Conversion
	/*
	   The process of converting a value from one data type to another is known as
	   type conversion. Converting from one datatype to another is a frequent task we
	   programmers do. Let's take few examples of converting variables from one data
	   type to another.
	*/

	stringToInt.StringToInt()
	stringToInt.StringToIntByParseInt()

	stringToFloat.StringToFloat()

	stringToBool.StringToBool()

	boolToString.BoolToStringUsingFormatBool()
	boolToString.BoolToStringUsingFmtSprintf()
	
	floatToString.FloatToStringUsingFmtSprintf()
	floatToString.FloatToStringUsingFormatFloat()

	intToString.IntToStringuUSingFormatString()
	intToString.IntToStringUsingFmtSprintf()

	intToSpecificInt.IntToSpecificInt()

	floatToSpecificFloat.FloatToSpecificFloat()

	intToFloat.IntToFloat()
}
