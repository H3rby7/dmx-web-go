package apiv1

type DMXValueForChannel struct {
	Channel int
	Value   int
}

type MultipleDMXValueForChannel struct {
	List []DMXValueForChannel
}
