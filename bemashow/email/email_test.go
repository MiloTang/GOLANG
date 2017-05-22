package email

import "testing"

func TestSendEmailWithAttachment(t *testing.T) {
	err := SendEmailWithAttachment("mainframe_fb@163.com", "!", "smtp.163.com:25", "168461237@qq.com", "测试附件")
	if err != nil {
		t.Fatal(err)
	}
}
