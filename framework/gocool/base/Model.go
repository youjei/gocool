package base

import (
	"fmt"
)

const (
	/**
	 * The name of the default scenario.
	 */
	SCENARIO_DEFAULT = "default"

	/**
	 * @event ModelEvent an event raised at the beginning of [[validate()]]. You may set
	 * [[ModelEvent::isValid]] to be false to stop the validation.
	 */
	EVENT_BEFORE_VALIDATE = "beforeValidate"

	/**
	 * @event Event an event raised at the end of [[validate()]]
	 */
	EVENT_AFTER_VALIDATE = "afterValidate"
)

type A struct {
	a int
}

func Model() {
	fmt.Printf("base/model")

}
