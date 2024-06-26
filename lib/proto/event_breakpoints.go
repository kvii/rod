// This file is generated by "./lib/proto/generate"

package proto

/*

EventBreakpoints

EventBreakpoints permits setting JavaScript breakpoints on operations and events
occurring in native code invoked from JavaScript. Once breakpoint is hit, it is
reported through Debugger domain, similarly to regular breakpoints being hit.

*/

// EventBreakpointsSetInstrumentationBreakpoint Sets breakpoint on particular native event.
type EventBreakpointsSetInstrumentationBreakpoint struct {
	// EventName Instrumentation name to stop on.
	EventName string `json:"eventName"`
}

// ProtoReq name.
func (m EventBreakpointsSetInstrumentationBreakpoint) ProtoReq() string {
	return "EventBreakpoints.setInstrumentationBreakpoint"
}

// Call sends the request.
func (m EventBreakpointsSetInstrumentationBreakpoint) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// EventBreakpointsRemoveInstrumentationBreakpoint Removes breakpoint on particular native event.
type EventBreakpointsRemoveInstrumentationBreakpoint struct {
	// EventName Instrumentation name to stop on.
	EventName string `json:"eventName"`
}

// ProtoReq name.
func (m EventBreakpointsRemoveInstrumentationBreakpoint) ProtoReq() string {
	return "EventBreakpoints.removeInstrumentationBreakpoint"
}

// Call sends the request.
func (m EventBreakpointsRemoveInstrumentationBreakpoint) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// EventBreakpointsDisable Removes all breakpoints.
type EventBreakpointsDisable struct{}

// ProtoReq name.
func (m EventBreakpointsDisable) ProtoReq() string { return "EventBreakpoints.disable" }

// Call sends the request.
func (m EventBreakpointsDisable) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}
