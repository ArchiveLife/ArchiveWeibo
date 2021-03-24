package provision

import (
	"github.com/ArchiveLife/core/adapter"
)

type WeiboServiceProvision struct {
}

func (p WeiboServiceProvision) ProvideServices() []adapter.ArchiveService {
	return []adapter.ArchiveService{
		createSingleUserWeiboService(),
	}
}
