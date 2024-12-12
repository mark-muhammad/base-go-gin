package rest

var handlers []RestHandler = []RestHandler{
	&AccountHandler{},
	&PersonHandler{},
}
