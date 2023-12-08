package models_chase

import (
	"time"

	models_scene "github.com/H3rby7/dmx-web-go/internal/model/scene"
	log "github.com/sirupsen/logrus"
)

// SceneRenderFunc defines the renderer used by chase as return value.
type SceneRenderFunc func(scene models_scene.Scene, fadeTimeMillis int64)

// ChangeBridgeStateFunc defines the function to turn the DMXBridge on/off
type ChangeBridgeStateFunc func(active bool)

// Internal representation of a Chase.
type Chase struct {
	// Name for this trigger - mus be unique
	Name string
	// The sequence of actions (chase) to take.
	Steps []Step
	// Delegate to render a step's scene
	renderDelegate SceneRenderFunc
	// Delegate to change bridge state
	bridgeDelegate ChangeBridgeStateFunc
	// Index of next step (in the chase)
	nextStep int
}

// NewChase constructs a Chase struct, given the name and steps to use.
func NewChase(name string, steps []Step) Chase {
	return Chase{
		Name:     name,
		Steps:    steps,
		nextStep: len(steps),
	}
}

// Run the chase continuing with the next step
func (c *Chase) RunFromStart(renderer SceneRenderFunc, bridge ChangeBridgeStateFunc) {
	log.WithField("chase", c.Name).Debugf("Starting chase from step 0")
	c.nextStep = 0
	c.renderDelegate = renderer
	c.bridgeDelegate = bridge
	go c.planNextStep()
}

// Take the next step and
//
// 1. Render it, using the render delegate
// 2. Increment next step
// 3. Call planNextStep
func (c *Chase) renderNextAndGoNext() {
	c.renderNextStep()
	log.WithField("chase", c.Name).Tracef("Incrementing next step to '%d'", c.nextStep)
	c.nextStep++
	c.planNextStep()
}

// Return the next step
//
// *Be careful to check with [hasNextStep] if we have one, first*
func (c *Chase) getNextStep() Step {
	return c.Steps[c.nextStep]
}

// Is there a next step to render
//
// compares the length of the chase to the nextStep
func (c *Chase) hasNextStep() bool {
	return (len(c.Steps) > c.nextStep)
}

// Get next step and render it
//
// Does nothing if chase is at its end or no render delegate is set
func (c *Chase) renderNextStep() {
	ll := log.WithField("chase", c.Name)
	ll.Tracef("Attempting to render next step (%d)", c.nextStep)

	if !c.hasNextStep() {
		ll.Infof("Reached end of chase")
		return
	}

	ll.Debugf("Rendering step '%d'", c.nextStep)
	nextStep := c.getNextStep()
	if c.renderDelegate == nil {
		ll.Warnf("Render delegate unset, skipping")
	} else {
		c.renderDelegate(nextStep.Scene, nextStep.FadeTimeMillis)
	}
	if c.bridgeDelegate == nil {
		ll.Warnf("Bridge delegate unset, skipping")
	} else {
		c.bridgeDelegate(nextStep.BridgeActive)
	}
}

// planNextStep waits the appropriate amount of time (delay)
// before calling [renderNextAndGoNext]
func (c *Chase) planNextStep() {
	ll := log.WithField("chase", c.Name)
	ll.Tracef("Planning next step (%d)", c.nextStep)

	if !c.hasNextStep() {
		ll.Infof("Reached end of chase")
		return
	}

	delay := c.getNextStep().DelayTimeMillis
	ll.Debugf("Planning next step (%d) in %d millis", c.nextStep, delay)
	time.Sleep(time.Duration(delay) * time.Millisecond)

	c.renderNextAndGoNext()
}
