package main

import (
	"encoding/json"
	"fmt"
)

package main

import (
"encoding/json"
"fmt"
"gorm.io/gorm"
)
type circle struct {
	conn *gorm.DB
}

func (g *circle) Add(p *PersonalInformation) error {
	resp := g.conn.Create(p)
	if err := resp.Error; err != nil {
		fmt.Println("创建圈子时失败：", err)
		return err
	}
	fmt.Println("创建圈子成功")
	return nil
}

func (g *circle) Get() (data []byte) {
	output := make([]*PersonalInformation, 0)
	resp := g.conn.Where("is_show = 1").Find(&output)
	if resp.Error != nil {
		fmt.Println("查找失败：", resp.Error)
		return
	}

	data, _ = json.Marshal(output)
	fmt.Printf("查询结果：%+v\n", string(data))
	return
}

func (g *circle) Delete(p *PersonalInformation) error {
	resp := g.conn.Model(p).Select("id", "is_show").Updates(p)
	if err := resp.Error; err != nil {
		fmt.Println("删除圈子时失败：", err)
		return err
	}
	fmt.Println("删除圈子成功")
	return nil
}
