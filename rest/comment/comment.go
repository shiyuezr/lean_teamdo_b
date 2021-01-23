package comment

import (
	"github.com/kfchen81/beego/vanilla"
)

type Comment struct {
	vanilla.RestResource
}

//用return的东西注册路由：vanilla/router.go中注册
func (this *Comment) Resource() string {
	return "comment.comment"
}

func (this *Comment) Get() {
	//id, _ := this.GetInt("id")
	////获取前端http请求的数据
	//bCtx := this.GetBusinessContext()
	//
	//////分页
	////page := vanilla.ExtractPageInfoFromRequest(this.Ctx)
	//////过滤原则
	////filters := common.ConvertToBeegoOrmFilter(this.GetFilters())
	//
	////使用bCtx构建CommentRepository对象
	//repository := comment.NewCommentRepository(bCtx)
	//Comment := repository.GetComment(id)
	//
	//////获得bCtx中携带的当前Corporation对象
	////corp := account.GetCorpFromContext(bCtx)
	////products, nextPageInfo := repository.GetEnabledProductsForCorp(corp, page, filters)
	//
	//fillService := comment.NewFillCommentService(bCtx)
	////fillService.Fill(products, vanilla.FillOption{
	////	"with_category": true,
	////	"with_tag":      true,
	////})
	//
	//encodeService := comment.NewEncodeCommentService(bCtx)
	//rows := encodeService.EncodeOne(Comment)
	//
	//response := vanilla.MakeResponse(vanilla.Map{
	//	"products": rows,
	//	"pageinfo": nextPageInfo.ToMap(),
	//})
	//this.ReturnJSON(response)
}

//func (this *Comment) Get(){
//	return
//}
