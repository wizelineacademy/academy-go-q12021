package presenter

type adminPresenter struct {
}

type AdminPresenter interface {
	ResponseLogs(logs []string) []string
}

func NewAdminPresenter() AdminPresenter {
	return &adminPresenter{}
}

func (ap *adminPresenter) ResponseLogs(logs []string) []string {
	return logs
}
