package logging

import (
	"github.com/topxeq/gowasm/api"
	. "github.com/topxeq/gowasm/internal/assemblyscript"
	"github.com/topxeq/gowasm/internal/logging"
)

func isProcFunction(fnd api.FunctionDefinition) bool {
	return fnd.ExportNames()[0] == AbortName
}

func isRandomFunction(fnd api.FunctionDefinition) bool {
	return fnd.ExportNames()[0] == SeedName
}

// IsInLogScope returns true if the current function is in any of the scopes.
func IsInLogScope(fnd api.FunctionDefinition, scopes logging.LogScopes) bool {
	if scopes.IsEnabled(logging.LogScopeProc) {
		if isProcFunction(fnd) {
			return true
		}
	}

	if scopes.IsEnabled(logging.LogScopeRandom) {
		if isRandomFunction(fnd) {
			return true
		}
	}

	return scopes == logging.LogScopeAll
}
