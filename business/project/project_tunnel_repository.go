package project

import (
	"context"
	"github.com/kfchen81/beego/vanilla"
)

type TunnelRepository struct {
	vanilla.RepositoryBase
}

func NewTunnelRepository(ctx context.Context) *TunnelRepository {
	repository := new(TunnelRepository)
	repository.Ctx = ctx
	return repository
}

func (this *TunnelRepository) GetTunnelsByProjectId(projectId int)  {

}
