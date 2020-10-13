package worker

import (
	"fmt"
	"learngo/18_crawler_distribution/concurrent_edition/engine"
	"learngo/18_crawler_distribution/concurrent_edition/fetcher"
	"learngo/18_crawler_distribution/concurrent_edition/parser/zhenai"
	"learngo/18_crawler_distribution/final_edition/config"
	"log"
)

type SerializedParser struct {
	Name string
	Args interface{}
}

// {"ParseCityList", nil}, {"ParseUserList", nil}, {"ParseUserInfo", nil}

type SerializedFetcher struct {
	Name string
	Args interface{}
}

// {"Fetch", nil}, {"IphoneFetch", nil}

type Request struct {
	Url     string
	Fetcher SerializedFetcher
	Parser  SerializedParser
}

type ParseResult struct {
	Requests []Request
	Items    []engine.Item
}

// define serialize and deSerialize
func SerializeRequest(r engine.Request) Request {
	nameFunc, argsFunc := r.Parser.Serialize()
	nameFetcher, argsFetcher := r.Fetcher.Serialize()
	return Request{
		Url: r.Url,
		Fetcher: SerializedFetcher{
			Name: nameFetcher,
			Args: argsFetcher,
		},
		Parser: SerializedParser{
			Name: nameFunc,
			Args: argsFunc,
		},
	}
}

func SerializeParseResult(r engine.ParseResult) ParseResult {
	result := ParseResult{}
	result.Items = r.Items

	for _, request := range r.Requests {
		re := SerializeRequest(request)
		result.Requests = append(result.Requests, re)
	}
	return result
}

func DeSerializeRequest(r Request) (engine.Request, error) {
	f, err := deSerializeFetcher(r.Fetcher)
	if err != nil {
		return engine.Request{}, err
	}

	p, err := deSerializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}

	return engine.Request{
		Url:     r.Url,
		Fetcher: f,
		Parser:  p,
	}, nil
}

func DeSerializeParseResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{}
	result.Items = r.Items

	for _, request := range r.Requests {
		re, err := DeSerializeRequest(request)
		if err != nil {
			log.Printf(" error deSerializing Request : %v", err)
			continue
		}
		result.Requests = append(result.Requests, re)
	}
	return result
}

func deSerializeFetcher(f SerializedFetcher) (engine.Fetcher, error) {
	switch f.Name {
	case config.NormalFetcher:
		return engine.NewFuncFetcher(fetcher.Fetch, config.NormalFetcher), nil
	case config.IphoneFetch:
		return engine.NewFuncFetcher(fetcher.Fetch, config.IphoneFetch), nil
	default:
		return nil, fmt.Errorf("unknown fetcher name: %s", f.Name)
	}
}

func deSerializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseCityList:
		return engine.NewFuncParser(zhenai.ParseCityList, config.ParseCityList), nil
	case config.ParseUserList:
		return engine.NewFuncParser(zhenai.ParseUserList, config.ParseUserList), nil
	case config.ParseUserInfo:
		return engine.NewFuncParser(zhenai.ParseUserInfo, config.ParseUserInfo), nil
	case config.NilParser:
		return engine.NewFuncParser(zhenai.NilParser, config.NilParser), nil
	default:
		return nil, fmt.Errorf("unknown parser name: %s", p.Name)
	}
}
