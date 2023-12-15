// Package models_chase contains the internal representations of [Chase] and [Step]
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
	// Name for this trigger - must be unique
	Name string
	// The sequence of actions (chase) to take.
	Steps []Step
	// Delegate to render a step's scene
	renderDelegate SceneRenderFunc
	// Delegate to change bridge state
	bridgeDelegate ChangeBridgeStateFunc
	// Index of next step (in the chase)
	nextStep int
	// Timer to trigger the next step
	timer *time.Timer
}

// NewChase constructs a Chase struct, given the name and steps to use.
func NewChase(name string, steps []Step) Chase {
	c := Chase{
		Name:     name,
		Steps:    steps,
		nextStep: len(steps), // So [isConcluded] returns true on a fresh chase
	}
	return c
}

// Run the chase continuing with the next step
func (c *Chase) RunFromStart(renderer SceneRenderFunc, bridge ChangeBridgeStateFunc) {
	log.WithField("chase", c.Name).Debugf("Resetting to step[0] and applying delegates for rendering and bridge")
	c.nextStep = 0
	c.renderDelegate = renderer
	c.bridgeDelegate = bridge
	c.planNextStep()
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
// *Be careful to check with [isConcluded] if we have one, first*
func (c *Chase) getNextStep() Step {
	return c.Steps[c.nextStep]
}

// Is this chase concluded/ended/not-started
//
// Checks if there is another step to render by comparing the length of the chase to the nextStep
func (c *Chase) isConcluded() bool {
	return (c.nextStep >= len(c.Steps))
}

// Get next step and render it
//
// Does nothing if chase is at its end or no render delegate is set
func (c *Chase) renderNextStep() {
	ll := log.WithField("chase", c.Name)
	ll.Tracef("Attempting to render next step (%d)", c.nextStep)

	if c.isConcluded() {
		ll.Infof("Reached end")
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

	if c.isConcluded() {
		ll.Infof("Reached end")
		return
	}

	delay := c.getNextStep().DelayTimeMillis
	if c.timer == nil {
		ll.Debugf("Creating new AfterFunc timer to render step[%d] in %d millis", c.nextStep, delay)
		c.timer = time.AfterFunc(time.Duration(delay)*time.Millisecond, c.renderNextAndGoNext)
	} else {
		ll.Debugf("Changing existing AfterFunc timer to render step[%d] in %d millis", c.nextStep, delay)
		c.timer.Stop()
		c.timer.Reset(time.Duration(delay) * time.Millisecond)
	}
}
