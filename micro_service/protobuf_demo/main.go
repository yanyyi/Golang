package main

import (
	"fmt"
	"github.com/golang/protobuf/protoc-gen-go/generator"
	"github.com/golang/protobuf/ptypes/wrappers"
	fieldmask_utils "github.com/mennanov/fieldmask-utils"
	"google.golang.org/genproto/protobuf/field_mask"
	"protobuf_demo/api"
)

// oneofDemo oneof示例
func oneofDemo() {
	//req1 := &api.NoticeReaderRequest{
	//	Msg: "李华考上清华大学",
	//	NoticeWay: &api.NoticeReaderRequest_Email{
	//		Email: "123456@qq.com",
	//	},
	//}

	req2 := &api.NoticeReaderRequest{
		Msg: "斯蒂芬库里率队赢得NBA总冠军",
		NoticeWay: &api.NoticeReaderRequest_Phone{
			Phone: "17867599676",
		},
	}

	//server
	//req.NoticeWay ???
	req := req2
	// 类型断言
	switch v := req.NoticeWay.(type) {
	case *api.NoticeReaderRequest_Email:
		noticeWithEmail(v)
	case *api.NoticeReaderRequest_Phone:
		noticeWithPhone(v)
	}
}

// 发送通知相关的功能函数
func noticeWithEmail(in *api.NoticeReaderRequest_Email) {
	fmt.Printf("notice reader by email(邮件):%v\n", in.Email)
}

func noticeWithPhone(in *api.NoticeReaderRequest_Phone) {
	fmt.Printf("notice reader by phone(手机):%v\n", in.Phone)
}

// wrapValueDemo 使用google/protobuf/wrappers.proto
func wrapValueDemo() {
	//client
	book := api.Book{
		Title:  "《Go语言入门与实战》",
		Author: "Google scientists team",
		Price:  &wrappers.Int64Value{Value: 89.00},
		Memo:   &wrappers.StringValue{Value: "打代码一定要多动手练习"},
	}
	if book.Price == nil { //没有给price赋值
		fmt.Println("没有设置price")
	} else {
		price := book.GetPrice().GetValue()
		fmt.Printf("%s的价格是:%d元\n", book.GetTitle(), price)
	}
	if book.Memo == nil {
		fmt.Println("没有设置memo")
	} else {
		memo := book.GetMemo().GetValue()
		fmt.Printf("%s的描述是:%s\n", book.GetTitle(), memo)
	}
}

// fieldMaskDemo 使用field_mask实现部分更新实例
func fieldMaskDemo() {
	// client
	paths := []string{"price", "info.b"} // 更新的字段信息
	req := &api.UpdateBookRequest{

		Op: "Stephen",
		Book: &api.Book{
			Price: &wrappers.Int64Value{Value: 67.00},
			Info: &api.Book_Info{
				B: "repeated bbbbb",
			},
		},
		UpdateMask: &field_mask.FieldMask{Paths: paths},
	}

	//server
	mask, _ := fieldmask_utils.MaskFromProtoFieldMask(req.UpdateMask, generator.CamelCase)
	var bookDst = make(map[string]interface{})
	//将数据读取到map[string]interface{}
	// fieldmask-utils支持读取到结构体等,更多用法可查看文档
	fieldmask_utils.StructToMap(mask, req.Book, bookDst)
	fmt.Printf("bookDst:%v\n", bookDst)

}

func main() {
	oneofDemo()
	wrapValueDemo()
	fieldMaskDemo()
}
