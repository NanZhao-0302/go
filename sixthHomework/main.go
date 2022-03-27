//功能需求 - 圈子

//发状态
//像朋友圈一样，发布心情、感想、幽默段子
//删除自己的状态
//逛圈子，看看整个圈子里边最新的状态
//圈子要求：
//发状态的内容包含： 1.发布时间 2.发布者 ID 3.发布者姓名 4.发布状态时的：年龄、身高、体重、体脂率
//发布状态后，每条都要落地到数据库
//删除数据时，标记圈子为不可见
//逛圈子时，看不见已经“删除”的状态

package apis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ClientInterface interface {
	ReadPostInformation() apis.Circle
	GetPersonId() uint32
}

type fakeCircleInterface struct {
	personId     uint32
	personName   string
	sex          string
	content      string
	atTimeHeight float32
	atTimeWeight float32
	atTimeAge    uint32
}

var _ ClientInterface = &fakeCircleInterface{}

func (f *fakeCircleInterface) ReadPostInformation() apis.Circle {
	cr := apis.Circle{
		PersonId:     f.personId,
		PersonName:   f.personName,
		Sex:          f.sex,
		Content:      f.content,
		AtTimeHeight: f.atTimeHeight,
		AtTimeWeight: f.atTimeWeight,
		AtTimeAge:    f.atTimeAge,
	}

	return cr
}
func (f fakeCircleInterface) GetPersonId() uint32 {
	return f.personId
}

type crClient struct {
	handRing ClientInterface
}

func (c crClient) post() {

	cr := c.handRing.ReadPostInformation()
	data, _ := json.Marshal(cr)
	r := bytes.NewReader(data)
	resp, err := http.Post("http://localhost:8081/post", "application/json", r)
	if err != nil {
		log.Println("登录失败", err)
		return
	}
	if resp.Body != nil {
		data, _ := ioutil.ReadAll(resp.Body)
		log.Println(string(data))
	}
}

func (c crClient) list() {
	resp, err := http.Get("http://localhost:8081/list")
	if err != nil {
		log.Println(err)
		return
	}
	if resp.Body != nil {
		data, _ := ioutil.ReadAll(resp.Body)
		log.Println("返回要求的列表", string(data))
	}
}

//删除圈子中的状态
func (c *crClient) delete(personId uint32) {
	url := fmt.Sprintf("http://localhost:8081/delete/%d", personId)
	log.Println("url", url)
	rep, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		log.Println("一个错误，删除失败:", err)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(rep)
	if err != nil {
		log.Println("一个错误，删除失败:", err)
		return
	}

	defer resp.Body.Close()

	if resp.Body != nil {
		data, _ := ioutil.ReadAll(resp.Body)
		log.Println(string(data))
	}
}

func main() {
	crCli := &crClient{handRing: &fakeCircleInterface{
		personId:     333,
		personName:   "ZhaoNan",
		sex:          "F",
		content:      "My post",
		atTimeHeight: 1.6,
		atTimeWeight: 50.0,
		atTimeAge:    18,
	},
	}
	//在我的圈子中发布状态
	crCli.post()
	//List top posts in circle
	crCli.list()
	//只看自己的朋友圈
	crCli.delete(crCli.handRing.GetPersonId())
	crCli.list()
}
