package julian

// ToJulianCentury converts a Julian day to century.
// 	jd: Julian day
// 	jc: Julian century
func ToJulianCentury(jd float64) (jc float64) {
	jc = jd * 31557600.0 / 3155695200.0
	return
}
