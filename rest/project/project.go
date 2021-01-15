package project

import (
	b_project "teamdo/business/project"

	"github.com/kfchen81/beego"
	"github.com/kfchen81/beego/vanilla"
)

type Project struct {
	vanilla.RestResource
}

func (this *Project) Resource() string {
	return "project.project"
}

func (this *Project) GetParamrters() map[string][]string {
	return map[string][]string{
		"GET": []string{
			"id:int",
		},
		"PUT": []string{
			"name:string",
			"user_ids:json-array",
		},
		"POST": []string{
			"id:int",
			"name:string",
			"introduction:string",
			"cover:string",
			"start_time:string",
			"user_ids:json-array",
		},
		"DELETE": []string{
			"id:int",
		},
	}
}

func (this *Project) Get() {
	id, err := this.GetInt("id") //从parameter获取参数
	if err != nil {
		beego.Error(err)
	}

	bCtx := this.GetBusinessContext()                    //获取上下文信息
	repository := b_project.NewProjectRepository(bCtx)   //实例化prject仓储
	project := repository.GetProject(id)                 //通过project的id获取project对象
	fillservice := b_project.NewFillProjectService(bCtx) //获取外键数据填充服务
	fillservice.Fill([]*b_project.Project{project}, vanilla.FillOption{
		//"with_task": true,
		"with_user": true,
	}) //获取项目携带的task和user

	encodeService := b_project.NewEncodeProjectService(bCtx) //实例化数据整形器

	respData := encodeService.Encode(project) //数据整形

	response := vanilla.MakeResponse(respData) //加入response

	this.ReturnJSON(response) //数据返回

}
