package http_utils

var CRLF = ([]byte)("\r\n")
var SPACE = []byte{' '}
var TAB = []byte{'\t'}
var CR = []byte{'\r'}
var LF = []byte{'\n'}
var NUL = []byte{0}

var RequestTypes = []string{"OPTIONS", "GET", "HEAD", "POST", "PUT", "DELETE", "TRACE", "CONNECT"}

// ToDo: Add all allowes Status Codes ?
// Check if the status code is valid
