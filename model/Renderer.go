package model

type Renderer interface {
	BeforeDraw()
	AfterDraw()
	Draw(*Cell)
}
