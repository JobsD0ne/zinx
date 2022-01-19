package ziface

type IRouter interface {
	PreHandle(requesst IRequest) //在处理conn业务之前的钩子方法
	Handle(request IRequest)     //处理conn业务的方法
	PostHandle(request IRequest) //处理conn业务方法之后的钩子方法
}
