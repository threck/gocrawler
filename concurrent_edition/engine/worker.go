package engine

import "log"

func Worker(re Request) (ParseResult, error) {
	//log.Printf("Fetching URL:%s", re.Url)
	bytes, err := re.Fetcher.Fetch(re.Url)
	if err != nil {
		log.Printf("Fetching URL:%s error: %s", re.Url, err)
		return ParseResult{}, err
	}

	//log.Printf("Parsing URL:%s", re.Url)
	result, err := re.Parser.Parse(bytes, re.Url)
	if err != nil {
		log.Printf("Parsing URL:%s error: %s", re.Url, err)
		return ParseResult{}, err
	}
	return result, nil
}
