package src

import "template/src/common"

func (app AppModule) CreateAppBuilder() *AppModule {
	return &app
}

func (app *AppModule) AddModule(module any) *AppModule {
	app.modules = append(app.modules, module.(common.IModule))
	return app
}

func (app *AppModule) Build() {
	for _, m := range app.modules {
		m.Bundle()
	}
}
