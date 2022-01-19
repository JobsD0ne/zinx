package znet

import "zinx/ziface"

type BaseRouter struct {
}

//这里之所以baseRouter的方法都为空
//是因为有的router不希望有prehandle或者posthandle
//所以router全部继承baseRouter的好处是，不需要实现preHandle和PostHandle也可以实例化
func (br *BaseRouter) PreHandle(rez ziface.IRequest)  {}
func (br *BaseRouter) Handle(rez ziface.IRequest)     {}
func (br *BaseRouter) PostHandle(rez ziface.IRequest) {}
