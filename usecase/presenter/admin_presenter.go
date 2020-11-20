package presenter

type AdminPresenter interface {
	ResponseLogs(logRecords []string) []string
}
