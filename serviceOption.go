package winsvc

import "golang.org/x/sys/windows/svc/mgr"

type ServiceOption func(*mgr.Config)

func DisplayName(displayName string) ServiceOption {
	return func(config *mgr.Config) {
		config.DisplayName = displayName
	}
}

func Description(description string) ServiceOption {
	return func(config *mgr.Config) {
		config.Description = description
	}
}

func OnBootStart() ServiceOption {
	return func(config *mgr.Config) {
		config.StartType = 0
	}
}

func OnSystemStart() ServiceOption {
	return func(config *mgr.Config) {
		config.StartType = 1
	}
}

func AutoStart() ServiceOption {
	return func(config *mgr.Config) {
		config.StartType = 2
	}
}

func AutoDelayStart() ServiceOption {
	return func(config *mgr.Config) {
		config.StartType = 2
		config.DelayedAutoStart = true
	}
}

func OnDemandStart() ServiceOption {
	return func(config *mgr.Config) {
		config.StartType = 3
	}
}

func DisabledStart() ServiceOption {
	return func(config *mgr.Config) {
		config.StartType = 4
	}
}

func Dependencies(serviceName ...string) ServiceOption {
	return func(config *mgr.Config) {
		for _, svcName := range serviceName {
			config.Dependencies = append(config.Dependencies, svcName)
		}
	}
}

type ServiceParamsOption func(param ...string) []string

func ServiceInstallParams(param ...string) ServiceParamsOption {
	return func(param ...string) []string {
		params := make([]string, 0)
		for _, p := range param {
			params = append(params, p)
		}
		return params
	}
}
