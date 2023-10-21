package dmxfader

import (
	apiv1 "github.com/H3rby7/dmx-web-go/internal/api/v1"
	dmxconn "github.com/H3rby7/dmx-web-go/internal/dmx"
	log "github.com/sirupsen/logrus"
)

const FADER_UPDATE_INTERVAL_MILLIS = 1000 / 25

type DMXFader struct {
	isFading       bool
	startState     apiv1.MultipleDMXValueForChannel
	targetState    apiv1.MultipleDMXValueForChannel
	fadeTimeMillis int
}

func NewDMXFader(target apiv1.MultipleDMXValueForChannel, fadeTimeMillis int) *DMXFader {
	f := &DMXFader{
		isFading:       false,
		targetState:    target,
		fadeTimeMillis: fadeTimeMillis,
	}
	f.setStartState()
	return f
}

// Start passing on any data that is read
func (f *DMXFader) Stop() {
	f.isFading = false
}

// Begin fading towards the targetState
func (f *DMXFader) Begin() {
	f.isFading = true
	steps := f.fadeTimeMillis / FADER_UPDATE_INTERVAL_MILLIS
	channelDeltaPerStep := make(map[int16]float32)
	for i := range f.targetState.List {
		channel := f.targetState.List[i].Channel
		target := f.targetState.List[i].Value
		start := f.startState.List[i].Value
		channelDeltaPerStep[channel] = float32(target-start) / float32(steps)
	}
	w := dmxconn.GetWriter()
	go func() {
		log.Infof("Starting to fade")
		for i := 0; i < steps; i++ {
			if !f.isFading {
				break
			}
			stage := w.GetStage()
			for _, e := range f.startState.List {
				ch := e.Channel
				nextVal := float32(stage[ch]) + channelDeltaPerStep[ch]
				w.Stage(ch, byte(nextVal))
			}
			w.Commit()
		}
		log.Infof("Fade finished")
	}()
}

// Set startState from writer's current stage
func (f *DMXFader) setStartState() {
	stage := dmxconn.GetWriter().GetStage()
	f.startState = apiv1.MultipleDMXValueForChannel{
		List: make([]apiv1.DMXValueForChannel, len(f.targetState.List)),
	}
	for i, v := range f.targetState.List {
		f.startState.List[i] = apiv1.DMXValueForChannel{
			Channel: v.Channel,
			Value:   stage[v.Channel],
		}
	}
}
