package syserrors

import (
	"github.com/MatthiasPetermann/sysexits"
)

// Base interface for custom errors with standardized exit codes
type CodedError interface {
	Error() string
	Code() int
}

// Usage error (EX_USAGE: 64)
type UsageError struct {
	Msg string
}

func (e *UsageError) Error() string {
	return e.Msg
}

func (e *UsageError) Code() int {
	return sysexits.EX_USAGE
}

// Data error (EX_DATAERR: 65)
type DataError struct {
	Msg string
}

func (e *DataError) Error() string {
	return e.Msg
}

func (e *DataError) Code() int {
	return sysexits.EX_DATAERR
}

// Input error (e.g., missing input file) (EX_NOINPUT: 66)
type NoInputError struct {
	Msg string
}

func (e *NoInputError) Error() string {
	return e.Msg
}

func (e *NoInputError) Code() int {
	return sysexits.EX_NOINPUT
}

// No user error (EX_NOUSER: 67)
type NoUserError struct {
	Msg string
}

func (e *NoUserError) Error() string {
	return e.Msg
}

func (e *NoUserError) Code() int {
	return sysexits.EX_NOUSER
}

// No host error (EX_NOHOST: 68)
type NoHostError struct {
	Msg string
}

func (e *NoHostError) Error() string {
	return e.Msg
}

func (e *NoHostError) Code() int {
	return sysexits.EX_NOHOST
}

// Remote error (e.g., API, network) (EX_UNAVAILABLE: 69)
type RemoteError struct {
	Msg string
}

func (e *RemoteError) Error() string {
	return e.Msg
}

func (e *RemoteError) Code() int {
	return sysexits.EX_UNAVAILABLE
}

// Internal software error (EX_SOFTWARE: 70)
type InternalSoftwareError struct {
	Msg string
}

func (e *InternalSoftwareError) Error() string {
	return e.Msg
}

func (e *InternalSoftwareError) Code() int {
	return sysexits.EX_SOFTWARE
}

// System error (EX_OSERR: 71)
type SystemError struct {
	Msg string
}

func (e *SystemError) Error() string {
	return e.Msg
}

func (e *SystemError) Code() int {
	return sysexits.EX_OSERR
}

// System file error (EX_OSFILE: 72)
type SystemFileError struct {
	Msg string
}

func (e *SystemFileError) Error() string {
	return e.Msg
}

func (e *SystemFileError) Code() int {
	return sysexits.EX_OSFILE
}

// Error: File cannot be created (EX_CANTCREAT: 73)
type FileCreationError struct {
	Msg string
}

func (e *FileCreationError) Error() string {
	return e.Msg
}

func (e *FileCreationError) Code() int {
	return sysexits.EX_CANTCREAT
}

// IO error (EX_IOERR: 74)
type IOError struct {
	Msg string
}

func (e *IOError) Error() string {
	return e.Msg
}

func (e *IOError) Code() int {
	return sysexits.EX_IOERR
}

// Temporary failure (EX_TEMPFAIL: 75)
type TemporaryError struct {
	Msg string
}

func (e *TemporaryError) Error() string {
	return e.Msg
}

func (e *TemporaryError) Code() int {
	return sysexits.EX_TEMPFAIL
}

// Protocol error (EX_PROTOCOL: 76)
type ProtocolError struct {
	Msg string
}

func (e *ProtocolError) Error() string {
	return e.Msg
}

func (e *ProtocolError) Code() int {
	return sysexits.EX_PROTOCOL
}

// Permission error (EX_NOPERM: 77)
type PermissionError struct {
	Msg string
}

func (e *PermissionError) Error() string {
	return e.Msg
}

func (e *PermissionError) Code() int {
	return sysexits.EX_NOPERM
}

// Configuration error (EX_CONFIG: 78)
type ConfigurationError struct {
	Msg string
}

func (e *ConfigurationError) Error() string {
	return e.Msg
}

func (e *ConfigurationError) Code() int {
	return sysexits.EX_CONFIG
}
