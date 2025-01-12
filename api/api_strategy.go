package api

import (
	"context"
	"github.com/filecoin-project/venus-wallet/core"
	"github.com/filecoin-project/venus-wallet/storage"
	"github.com/filecoin-project/venus-wallet/storage/strategy"
)

var _ strategy.ILocalStrategy = &StrategyAPIAdapter{}

type StrategyAPIAdapter struct {
	Internal struct {
		NewMsgTypeTemplate     func(ctx context.Context, name string, codes []int) error                                                  `perm:"admin" local:"required"`
		NewMethodTemplate      func(ctx context.Context, name string, methods []string) error                                             `perm:"admin" local:"required"`
		NewKeyBindCustom       func(ctx context.Context, name string, address core.Address, codes []int, methods []core.MethodName) error `perm:"admin" local:"required"`
		NewKeyBindFromTemplate func(ctx context.Context, name string, address core.Address, mttName, mtName string) error                 `perm:"admin" local:"required"`
		NewGroup               func(ctx context.Context, name string, keyBindNames []string) error                                        `perm:"admin" local:"required"`
		NewStToken             func(ctx context.Context, groupName string) (token string, err error)                                      `perm:"admin" local:"required"`

		GetMsgTypeTemplate      func(ctx context.Context, name string) (*storage.MsgTypeTemplate, error)    `perm:"admin" local:"required"`
		GetMethodTemplateByName func(ctx context.Context, name string) (*storage.MethodTemplate, error)     `perm:"admin" local:"required"`
		GetKeyBindByName        func(ctx context.Context, name string) (*storage.KeyBind, error)            `perm:"admin" local:"required"`
		GetKeyBinds             func(ctx context.Context, address core.Address) ([]*storage.KeyBind, error) `perm:"admin" local:"required"`
		GetGroupByName          func(ctx context.Context, name string) (*storage.Group, error)              `perm:"admin" local:"required"`
		GetWalletTokensByGroup  func(ctx context.Context, groupName string) ([]string, error)               `perm:"admin" local:"required"`
		GetWalletTokenInfo      func(ctx context.Context, token string) (*storage.GroupAuth, error)         `perm:"admin" local:"required"`

		ListGroups           func(ctx context.Context, fromIndex, toIndex int) ([]*storage.Group, error)           `perm:"admin" local:"required"`
		ListKeyBinds         func(ctx context.Context, fromIndex, toIndex int) ([]*storage.KeyBind, error)         `perm:"admin" local:"required"`
		ListMethodTemplates  func(ctx context.Context, fromIndex, toIndex int) ([]*storage.MethodTemplate, error)  `perm:"admin" local:"required"`
		ListMsgTypeTemplates func(ctx context.Context, fromIndex, toIndex int) ([]*storage.MsgTypeTemplate, error) `perm:"admin" local:"required"`

		AddMsgTypeIntoKeyBind    func(ctx context.Context, name string, codes []int) (*storage.KeyBind, error)      `perm:"admin" local:"required"`
		AddMethodIntoKeyBind     func(ctx context.Context, name string, methods []string) (*storage.KeyBind, error) `perm:"admin" local:"required"`
		RemoveMsgTypeFromKeyBind func(ctx context.Context, name string, codes []int) (*storage.KeyBind, error)      `perm:"admin" local:"required"`
		RemoveMethodFromKeyBind  func(ctx context.Context, name string, methods []string) (*storage.KeyBind, error) `perm:"admin" local:"required"`

		RemoveStToken          func(ctx context.Context, token string) error                                                  `perm:"admin" local:"required"`
		RemoveKeyBind          func(ctx context.Context, name string) error                                                   `perm:"admin" local:"required"`
		RemoveKeyBindByAddress func(ctx context.Context, address core.Address) (int64, error)                                 `perm:"admin" local:"required"`
		RemoveGroup            func(ctx context.Context, name string) error                                                   `perm:"admin" local:"required"`
		RemoveMethodTemplate   func(ctx context.Context, name string) error                                                   `perm:"admin" local:"required"`
		RemoveMsgTypeTemplate  func(ctx context.Context, name string) error                                                   `perm:"admin" local:"required"`
		ScopeWallet            func(ctx context.Context) (*core.AddressScope, error)                                          `perm:"admin" local:"required"`
		ContainWallet          func(ctx context.Context, address core.Address) bool                                           `perm:"admin" local:"required"`
		Verify                 func(ctx context.Context, address core.Address, msgType core.MsgType, msg *core.Message) error `perm:"admin" local:"required"`
	}
}

func (o *StrategyAPIAdapter) NewStToken(ctx context.Context, groupName string) (token string, err error) {
	return o.Internal.NewStToken(ctx, groupName)
}

func (o *StrategyAPIAdapter) GetWalletTokensByGroup(ctx context.Context, groupName string) ([]string, error) {
	return o.Internal.GetWalletTokensByGroup(ctx, groupName)
}

func (o *StrategyAPIAdapter) NewMsgTypeTemplate(ctx context.Context, name string, codes []int) error {
	return o.Internal.NewMsgTypeTemplate(ctx, name, codes)
}
func (o *StrategyAPIAdapter) NewMethodTemplate(ctx context.Context, name string, methods []string) error {
	return o.Internal.NewMethodTemplate(ctx, name, methods)
}
func (o *StrategyAPIAdapter) NewKeyBindCustom(ctx context.Context, name string, address core.Address, codes []int, methods []core.MethodName) error {
	return o.Internal.NewKeyBindCustom(ctx, name, address, codes, methods)
}
func (o *StrategyAPIAdapter) NewKeyBindFromTemplate(ctx context.Context, name string, address core.Address, mttName, mtName string) error {
	return o.Internal.NewKeyBindFromTemplate(ctx, name, address, mttName, mtName)
}
func (o *StrategyAPIAdapter) NewGroup(ctx context.Context, name string, keyBindNames []string) error {
	return o.Internal.NewGroup(ctx, name, keyBindNames)
}

func (o *StrategyAPIAdapter) GetMsgTypeTemplate(ctx context.Context, name string) (*storage.MsgTypeTemplate, error) {
	return o.Internal.GetMsgTypeTemplate(ctx, name)
}
func (o *StrategyAPIAdapter) GetMethodTemplateByName(ctx context.Context, name string) (*storage.MethodTemplate, error) {
	return o.Internal.GetMethodTemplateByName(ctx, name)
}
func (o *StrategyAPIAdapter) GetKeyBindByName(ctx context.Context, name string) (*storage.KeyBind, error) {
	return o.Internal.GetKeyBindByName(ctx, name)
}
func (o *StrategyAPIAdapter) GetKeyBinds(ctx context.Context, address core.Address) ([]*storage.KeyBind, error) {
	return o.Internal.GetKeyBinds(ctx, address)
}
func (o *StrategyAPIAdapter) GetGroupByName(ctx context.Context, name string) (*storage.Group, error) {
	return o.Internal.GetGroupByName(ctx, name)
}
func (o *StrategyAPIAdapter) GetWalletTokenInfo(ctx context.Context, token string) (*storage.GroupAuth, error) {
	return o.Internal.GetWalletTokenInfo(ctx, token)
}
func (o *StrategyAPIAdapter) ListGroups(ctx context.Context, fromIndex, toIndex int) ([]*storage.Group, error) {
	return o.Internal.ListGroups(ctx, fromIndex, toIndex)
}
func (o *StrategyAPIAdapter) ListKeyBinds(ctx context.Context, fromIndex, toIndex int) ([]*storage.KeyBind, error) {
	return o.Internal.ListKeyBinds(ctx, fromIndex, toIndex)
}
func (o *StrategyAPIAdapter) ListMethodTemplates(ctx context.Context, fromIndex, toIndex int) ([]*storage.MethodTemplate, error) {
	return o.Internal.ListMethodTemplates(ctx, fromIndex, toIndex)
}
func (o *StrategyAPIAdapter) ListMsgTypeTemplates(ctx context.Context, fromIndex, toIndex int) ([]*storage.MsgTypeTemplate, error) {
	return o.Internal.ListMsgTypeTemplates(ctx, fromIndex, toIndex)
}

func (o *StrategyAPIAdapter) AddMsgTypeIntoKeyBind(ctx context.Context, name string, codes []int) (*storage.KeyBind, error) {
	return o.Internal.AddMsgTypeIntoKeyBind(ctx, name, codes)
}
func (o *StrategyAPIAdapter) AddMethodIntoKeyBind(ctx context.Context, name string, methods []string) (*storage.KeyBind, error) {
	return o.Internal.AddMethodIntoKeyBind(ctx, name, methods)
}
func (o *StrategyAPIAdapter) RemoveMsgTypeFromKeyBind(ctx context.Context, name string, codes []int) (*storage.KeyBind, error) {
	return o.Internal.RemoveMsgTypeFromKeyBind(ctx, name, codes)
}
func (o *StrategyAPIAdapter) RemoveMethodFromKeyBind(ctx context.Context, name string, methods []string) (*storage.KeyBind, error) {
	return o.Internal.RemoveMethodFromKeyBind(ctx, name, methods)
}

func (o *StrategyAPIAdapter) RemoveKeyBind(ctx context.Context, name string) error {
	return o.Internal.RemoveKeyBind(ctx, name)
}
func (o *StrategyAPIAdapter) RemoveKeyBindByAddress(ctx context.Context, address core.Address) (int64, error) {
	return o.Internal.RemoveKeyBindByAddress(ctx, address)
}
func (o *StrategyAPIAdapter) RemoveGroup(ctx context.Context, name string) error {
	return o.Internal.RemoveGroup(ctx, name)
}
func (o *StrategyAPIAdapter) RemoveMethodTemplate(ctx context.Context, name string) error {
	return o.Internal.RemoveMethodTemplate(ctx, name)
}
func (o *StrategyAPIAdapter) RemoveMsgTypeTemplate(ctx context.Context, name string) error {
	return o.Internal.RemoveMsgTypeTemplate(ctx, name)
}
func (o *StrategyAPIAdapter) RemoveStToken(ctx context.Context, token string) error {
	return o.Internal.RemoveStToken(ctx, token)
}
func (o *StrategyAPIAdapter) ScopeWallet(ctx context.Context) (*core.AddressScope, error) {
	return o.Internal.ScopeWallet(ctx)
}
func (o *StrategyAPIAdapter) Verify(ctx context.Context, address core.Address, msgType core.MsgType, msg *core.Message) error {
	return o.Internal.Verify(ctx, address, msgType, msg)
}
func (o *StrategyAPIAdapter) ContainWallet(ctx context.Context, address core.Address) bool {
	return o.Internal.ContainWallet(ctx, address)
}
