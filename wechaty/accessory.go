package wechaty

import (
	wechatypuppet "github.com/kirito-er/go-wechaty/wechaty-puppet"
	_interface "github.com/kirito-er/go-wechaty/wechaty/interface"
)

// Accessory ...
type Accessory struct {
	puppet  wechatypuppet.IPuppetAbstract
	wechaty *Wechaty
}

// GetPuppet ...
func (a *Accessory) GetPuppet() wechatypuppet.IPuppetAbstract {
	return a.puppet
}

// GetWechaty ...
func (a *Accessory) GetWechaty() _interface.IWechaty {
	return a.wechaty
}
