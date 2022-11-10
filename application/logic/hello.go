package logic

import (
	"fmt"
	"go_project_demo/common"
	"go_project_demo/model"
)

const helloScrmRobotKey = "76da0600-c423-431a-bcbf-8a2e0258f237"

// SayHello 示例业务逻辑
func SayHello(firstName, lastName string) (rsp *common.Response, err error) {
	rsp = common.NewOKResponse()
	err = model.SendScrmMessage(helloScrmRobotKey, &model.ScrmRobotMessage{
		MsgType: "text",
		Text: &model.ScrmRobotMessageTextContent{
			Content: fmt.Sprintf(
				"【SayHello】: Hello from %s %s",
				firstName, lastName,
			),
			MentionedList:       []string{"@all"},
			MentionedMobileList: []string{"13919154925"},
		},
	})

	if err != nil {
		rsp.Code = common.ResponseCodeInternalErr
		rsp.Message = "发送企微robot消息失败"
		err = fmt.Errorf("model.SendScrmMessage: %w", err)
		return
	}

	return
}
