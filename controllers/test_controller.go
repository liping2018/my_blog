package controllers

type TestController struct {
	BaseController
}

func (c *TestController) URLMapping() {
	c.Mapping("test", c.TestRequestMessage)
}

//@router test [get]
func (c *TestController) TestRequestMessage() {
	c.ok()
}
