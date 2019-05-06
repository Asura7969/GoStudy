package engine

type Request struct {
	Url       string
	PaserFunc func([]byte) PaserResult
}

type PaserResult struct {
	Requests []Request
	Items    []interface{}
}

func NilParser([]byte) PaserResult {
	return PaserResult{}
}
