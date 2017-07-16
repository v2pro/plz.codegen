package gen

import (
	"plugin"
	"sync"
)

var loadPluginMutex = &sync.Mutex{}
var plugins = []*plugin.Plugin{}

// LoadPlugin in the run time, from .so file
func LoadPlugin(soFileName string) {
	loadPluginMutex.Lock()
	defer loadPluginMutex.Unlock()
	thePlugin, err := plugin.Open(soFileName)
	if err != nil {
		panic("failed to load generated plugin: " + err.Error())
	}
	plugins = append(plugins, thePlugin)
}

func lookupFuncFromPlugins(funcName string) interface{} {
	for _, thePlugin := range plugins {
		symbol, err := thePlugin.Lookup(funcName)
		logger.Debug("lookup func", "funcName", funcName, "plugin", thePlugin, "err", err)
		if err == nil {
			return symbol
		}
	}
	return nil
}
