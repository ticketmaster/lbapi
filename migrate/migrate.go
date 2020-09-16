package migrate

type Data struct {
	Response Response
}

func New() *Data {
	return &Data{}
}

// validate - ensures that the load balancer object meets the minimum requirements for submission.
