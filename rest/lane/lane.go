package lane

import (
	"github.com/kfchen81/beego/vanilla"
	b_lane "teamdo/business/lane"
)
type Lane struct {
	vanilla.RestResource
}


func (this *Lane) Resource() string {
	return "lane.lane"
}

func (this *Lane) GetParameters() map[string][]string {
	return map[string][]string{
		"GET": []string{},
		"PUT": []string{
			"lane_name:string",
			"project_id:int",
			"lane_sort_id:int",
		},
		"POST": []string{
			"lane_id:int",
			"lane_sort_id:int",
			"lane_name:string",
		},
		"DELETE": []string{
			"lane_id:int",
		},
	}
}

func (this *Lane)Put()  {
	bCtx:=this.GetBusinessContext()
	laneName:=this.GetString("lane_name")
	projectId,_:=this.GetInt("project_id")
	laneSortId,_:=this.GetInt("lane_sort_id")
	lane:= b_lane.NewLane(bCtx,laneName,projectId,laneSortId)
	response := vanilla.MakeResponse(vanilla.Map{
		"id": lane.Id,
	})
	this.ReturnJSON(response)
}

func (this *Lane)Post()  {
	bCtx:=this.GetBusinessContext()
	laneId,_:=this.GetInt("lane_id")
	laneSortId,_:=this.GetInt("lane_sort_id")
	laneName:=this.GetString("lane_name")
	lane:= b_lane.NewLaneRepository(bCtx).GetLane(laneId)
	err:= lane.Update(laneId,laneName,laneSortId)
	if err!=nil {
		panic(err)
	}
	response := vanilla.MakeResponse(vanilla.Map{})
	this.ReturnJSON(response)
}

func (this *Lane)Delete()  {
	bCtx:=this.GetBusinessContext()
	laneId,_:=this.GetInt("lane_id")
	lane:= b_lane.NewLaneRepository(bCtx).GetLane(laneId)
	err:= lane.Delete(laneId)
	if err!=nil {
		panic(err)
	}
	response := vanilla.MakeResponse(vanilla.Map{})
	this.ReturnJSON(response)
}

