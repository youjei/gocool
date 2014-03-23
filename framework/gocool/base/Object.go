package base

import (
   "io"
)

type Object struct {
	ArrayConfig bool=false
	FileConfig bool=false
	XMLConfig bool=false
	JasonConfig bool=false
}

/**
 * Configures an object with the initial property values.
 * @param object $object the object to be configured
 * @param array $properties the property initial values given in terms of name-value pairs.
 * @return object the object itself
 */
func (o *Object) Configure() {
  
}
