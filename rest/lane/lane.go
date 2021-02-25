package lane

import (
	"github.com/kfchen81/beego/vanilla"
	b_lane "teamdo/business/lane"
	b_project "teamdo/business/project"
)

type Lane struct {
	vanilla.RestResource
}

func (this *Lane) Resource() string {
	return "lane.lane"
}

func (this *Lane) GetParameters() map[string][]string {
	return map[string][]string{
		"GET": []string{"id:int", "?with_options:json"},
		"PUT": []string{
			"name:string", "pid:int",
		},
		"POST": []string{
			"id:int",
			"name:string",
			"pid:int",
			"sort:int",
		},
		"DELETE": []string{"id:int", "pid:int"},
	}
}

func (this *Lane) Get() {
	id, _ := this.GetInt("id")
	bCtx := this.GetBusinessContext()

	repository := b_lane.NewLaneRepository(bCtx)
	lane := repository.GetLaneById(id)
	if lane == nil {
		panic(vanilla.NewBusinessError("lane_not_exist", "泳道不存在"))
	}

	encodeService := b_lane.NewEncodeLaneService(bCtx)
	respData := encodeService.Encode(lane)
	response := vanilla.MakeResponse(respData)
	this.ReturnJSON(response)
}

func (this *Lane) Put() {
	name := this.GetString("name")
	pid, _ := this.GetInt("pid")
	bCtx := this.GetBusinessContext()
	project := b_project.NewProjectRepository(bCtx).GetProjectById(pid)
	project.AuthorityVerify()

	lane := b_lane.NewLane(bCtx, name, pid)
	lane.Update(lane.Name,lane.Id)
	response := vanilla.MakeResponse(vanilla.Map{
		"id": lane.Id,
	})
	this.ReturnJSON(response)
}

func (this *Lane) Delete() {
	id, _ := this.GetInt("id")
	pid, _ := this.GetInt("pid")
	bCtx := this.GetBusinessContext()

	project := b_project.NewProjectRepository(bCtx).GetProjectById(pid)
	project.AuthorityVerify()
	lane := b_lane.NewLaneRepository(bCtx).GetLaneById(id)
	lane.Delete()

	response := vanilla.MakeResponse(vanilla.Map{
		"id": id,
	})
	this.ReturnJSON(response)
}

func (this *Lane) Post() {
	id, _ := this.GetInt("id")
	pid, _ := this.GetInt("pid")
	sort, _ := this.GetInt("sort")
	name := this.GetString("name")
	bCtx := this.GetBusinessContext()

	project := b_project.NewProjectRepository(bCtx).GetProjectById(pid)
	project.AuthorityVerify()

	repository := b_lane.NewLaneRepository(bCtx)
	lane := repository.GetLaneById(id)
	lane.Update(name, sort)

	response := vanilla.MakeResponse(vanilla.Map{})
	this.ReturnJSON(response)
}
