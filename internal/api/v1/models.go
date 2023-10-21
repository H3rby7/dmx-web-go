package apiv1

type DMXValueForChannel struct {
	Channel int16
	Value   byte
}

type MultipleDMXValueForChannel struct {
	List []DMXValueForChannel
}

type MultipleDMXValueForChannelWithFade struct {
	FadeTimeMillis int
	Scene          MultipleDMXValueForChannel
}
