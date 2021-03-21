package calendatext

func Match(s string, baseDate Date, targetDate Date, baseEnabled bool) (bool, error) {
	cal := NewCalendar(baseDate, targetDate, baseEnabled)
	if err := cal.ParseText(s); err != nil {
		return false, err
	}
	return cal.Dates().Match(&targetDate), nil
}
