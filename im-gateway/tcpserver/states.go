package tcpserver

import goreactor "github.com/markity/go-reactor"

type ConnContext struct {
	ApiIMAuthRoute interface{}
	ApiIMRelay     interface{}
	State          ConnState
}

type ConnState struct {
	State StateEnum
	// 所有状态下都有的变量
	CommState CommState
}

type StateEnum int

const (
	StateEnumDefault = 0 // useless

	/* 状态说明
	用户在连上长连接网关后的状态是StateEnumUnLogin
	在此状态下, 如果在10s内没有进入登陆状态, 则断开连接
	状态转移:
	null (连上长连接网关)-> StateEnumUnLogin
	StateEnumUnLogin (登陆ok)-> StateEnumLogin
	StateEnumLogin (注销ok)-> StateEnumUnLogin
	*/
	StateEnumUnLogin = 1
	StateEnumLogin   = 2
)

type CommState struct {
	// 发送心跳包的timer id
	SendHeartbeatTimerId int
	//
	RecvHeartbeatTimerId int
}

func InitContext(conn goreactor.TCPConnection) {
	conn.SetContext("conn_context", &ConnContext{
		ApiIMAuthRoute: nil,
		ApiIMRelay:     nil,
		State: ConnState{
			State: StateEnumUnLogin,
		},
	})

}

func GetConnContextFromConn(conn goreactor.TCPConnection) *ConnContext {
	return conn.MustGetContext("conn_context").(*ConnContext)
}
