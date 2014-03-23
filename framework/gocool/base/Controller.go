package base

import ()

const (
	/**
	 * @event ActionEvent an event raised right before executing a controller action.
	 * You may set [[ActionEvent::isValid]] to be false to cancel the action execution.
	 */
	EVENT_BEFORE_ACTION = "beforeAction"
	/**
	 * @event ActionEvent an event raised right after executing a controller action.
	 */
	EVENT_AFTER_ACTION = "afterAction"
)

var (
	/**
	 * @var string the ID of this controller.
	 */
	Id int
	/**
	 * @var Module $module the module that this controller belongs to.
	 */
	Module string
	/**
	 * @var string the ID of the action that is used when the action ID is not specified
	 * in the request. Defaults to 'index'.
	 */
	DefaultAction string = "index"
	/**
	 * @var string|boolean the name of the layout to be applied to this controller's views.
	 * This property mainly affects the behavior of [[render()]].
	 * Defaults to null, meaning the actual layout value should inherit that from [[module]]'s layout value.
	 * If false, no layout will be applied.
	 */
	Layout string
	/**
	 * @var Action the action that is currently being executed. This property will be set
	 * by [[run()]] when it is called by [[Application]] to run an action.
	 */
	action string
	/**
	 * @var View the view object that can be used to render views or view files.
	 */
	_view string
)
