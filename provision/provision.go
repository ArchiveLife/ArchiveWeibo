package provision

import (
	"github.com/ArchiveLife/model/adapter"
)

type WeiboServiceProvision struct {
}

func (p WeiboServiceProvision) ProvideServices() []adapter.ArchiveService {
	return []adapter.ArchiveService{
		createSingleUserWeiboService(),
	}
}
