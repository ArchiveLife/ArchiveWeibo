package adapter

type ArchiveServiceProvision interface {
	ProvideServices() []ArchiveService
}
