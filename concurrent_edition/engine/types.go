package engine

type ParserFunc func(contents []byte, url string) (ParseResult, error)
type FetcherFunc func(url string) ([]byte, error)

type Parser interface {
	Parse(contents []byte, url string) (ParseResult, error)
	Serialize() (name string, args interface{})
}

type Fetcher interface {
	Fetch(url string) ([]byte, error)
	Serialize() (name string, args interface{})
}

type Request struct {
	Url     string
	Fetcher Fetcher
	Parser  Parser
}

//type Request struct {
//	Url        string
//	FetchFunc  func(string) ([]byte, error)
//	ParserFunc func([]byte, string) (ParseResult, error)
//}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Index   string
	Id      string
	Url     string
	Payload interface{}
}

//type NilParser struct{}
//
//func (n NilParser) Parse([]byte, string) (ParseResult, error) {
//	return ParseResult{}, nil
//}
//
//func (n NilParser) Serialize() (name string, args interface{}) {
//	return "NilParser", nil
//}

type FuncParser struct {
	parser ParserFunc
	name   string
}

func (f *FuncParser) Parse(contents []byte, url string) (ParseResult, error) {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}

type FuncFetcher struct {
	fetcher FetcherFunc
	name    string
}

func (f *FuncFetcher) Fetch(url string) ([]byte, error) {
	return f.fetcher(url)
}

func (f *FuncFetcher) Serialize() (name string, args interface{}) {
	return f.name, nil
}

func NewFuncFetcher(f FetcherFunc, name string) *FuncFetcher {
	return &FuncFetcher{
		fetcher: f,
		name:    name,
	}
}
