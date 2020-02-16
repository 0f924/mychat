package packet

// 响应数据包

// type: 创建群组
type CreateGroupResponsePacket struct {
	IsSuccess    bool
	GroupId      string
	UserNameList []string
}

func (packet CreateGroupResponsePacket) GetType() byte {
	return CREATE_GROUP_RESPONSE
}

// 创建群组 end

// type: 发送群组信息（待实现）
type GroupMessageResponsePacket struct {
	FromGroupId string
	// FromUser Session
	Message string
}

func (packet GroupMessageResponsePacket) GetType() byte {
	return GROUP_MESSAGE_RESPONSE
}

// 发送群组信息 end

// type: 心跳响应包
type HeartBeatResponsePacket struct {
}

func (packet HeartBeatResponsePacket) GetType() byte {
	return HEARTBEAT_RESPONSE
}

// 心跳响应包 end

// type: 加入群组
type JoinGroupResponsePacket struct {
	IsSuccess bool
	Reason    string
	GroupId   string
}

func (packet JoinGroupResponsePacket) GetType() byte {
	return JOIN_GROUP_RESPONSE
}

// 加入群组 end

// type: 列出群组成员（待实现）
type ListGroupMembersResponsePacket struct {
	GroupId string
	// SessionList []Session
}

func (packet ListGroupMembersResponsePacket) GetType() byte {
	return LIST_GROUP_MEMBERS_RESPONSE
}

// 列出群组成员 end

// type: 登录验证
type LoginResponsePacket struct {
	IsSuccess bool
	Reason    string
	UserId    uint64
	UserName  string
}

func (packet LoginResponsePacket) GetType() byte {
	return LOGIN_RESPONSE
}

// 登录验证 end

// type: 注销响应
type LogoutResponsePacket struct {
	IsSuccess bool
	Reason    string
}

func (packet LogoutResponsePacket) GetType() byte {
	return LOGOUT_RESPONSE
}

// 注销响应 end

// type: 单发信息
type MessageResponsePacket struct {
	FromUserId   string
	FromUserName string
	Message      string
}

func (packet MessageResponsePacket) GetType() byte {
	return MESSAGE_RESPONSE
}

// 单发信息 end

// type: 退出群组
type QuitGroupResponsePacket struct {
	IsSuccess bool
	Reason    string
	GroupId   string
}

func (packet QuitGroupResponsePacket) GetType() byte {
	return QUIT_GTOUP_RESPONSE
}

// 退出群组 end
