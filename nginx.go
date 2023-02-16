package main

type nginx struct{
	application server
	maxAllowedRequest int
	rateLimiter map[string]int
}

func newnginxserver() *nginx {
	nginx := &nginx{
		application : &application{},
		maxAllowedRequest: 2,
		rateLimiter: make(map[string]int),
	}
	return nginx
}

func (ng *nginx) handleRequest(url, method string) (int, string){
	allowed := ng.checkRateLimiting(url)
	if !allowed{
		return 403, "Not allowed"
	}
	return ng.application.handleRequest(url, method)
}

func (ng *nginx) checkRateLimiting(url string) bool{
	if ng.rateLimiter[url] == 0{
		ng.rateLimiter[url] = 1
	}
	if ng.rateLimiter[url] > ng.maxAllowedRequest{
		return false
	}
	ng.rateLimiter[url] += 1
	return true
}