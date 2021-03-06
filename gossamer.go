package gossamer

import "errors"

var ERR_QUERYOPTION_INVALID_LENGTH = errors.New("Invalid Query Option. Query Options must have a key and value.")
var ERR_QUERYOPTION_BLANK = errors.New("Query option string blank")
var ERR_QUERYOPTION_INVALID_VALUE = errors.New("Invalid Query Option value")
var ERR_INVALID_ENTITY = errors.New("Invalid Entity")

// Globals
var GLOB_ENV_HOST = "localhost:8000"
