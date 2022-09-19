package client_go

import "github.com/wuqinqiang/easycar/proto"

type (
	Branch struct {
		uri          string
		data, header []byte
		action       Action
		level        Level
		timeout      int64
		protocol     Protocol
	}
)

func NewBranch(uri string, action Action) *Branch {
	b := &Branch{
		uri:    uri,
		action: action,
		level:  1,
	}
	return b
}

func (branch *Branch) SetLevel(level Level) {
	if level > 1 {
		branch.level = level
	}
}

func (branch *Branch) SetProtocol(protocol Protocol) *Branch {
	branch.protocol = protocol
	return branch
}

func (branch *Branch) SetData(data []byte) *Branch {
	branch.data = data
	return branch
}

func (branch *Branch) SetHeader(header []byte) *Branch {
	branch.header = header
	return branch
}

func (branch *Branch) SetTimeout(timeout int64) *Branch {
	branch.timeout = timeout
	return branch
}

func (branch *Branch) Convert() *proto.RegisterReq_Branch {
	var (
		req proto.RegisterReq_Branch
	)
	req.Uri = branch.uri
	req.ReqData = string(branch.data)
	req.ReqHeader = string(branch.header)
	req.Protocol = string(branch.protocol)
	req.Timeout = int32(branch.timeout)
	req.Action = ConvertBranchActionToGrpc(branch.action)
	req.Level = int32(branch.level)
	return &req
}
