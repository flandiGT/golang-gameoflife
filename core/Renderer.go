package core

type Renderer interface {
	BeforeDraw()
	AfterDraw()
	Draw(*Cell)
}
