
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
//Code generated by mockery v1.0.0. 不要编辑。
package mocks

import mock "github.com/stretchr/testify/mock"
import policies "github.com/hyperledger/fabric/common/policies"

//ChannelPolicyManagergetter是ChannelPolicyManagergetter类型的自动生成的模拟类型
type ChannelPolicyManagerGetter struct {
	mock.Mock
}

//管理器提供具有给定字段的模拟函数：channelid
func (_m *ChannelPolicyManagerGetter) Manager(channelID string) (policies.Manager, bool) {
	ret := _m.Called(channelID)

	var r0 policies.Manager
	if rf, ok := ret.Get(0).(func(string) policies.Manager); ok {
		r0 = rf(channelID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(policies.Manager)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(channelID)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}
