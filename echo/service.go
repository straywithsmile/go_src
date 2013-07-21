package generator

type ServicePlugin struct {
	*Generator
}

func (p *ServicePlugin) Name() string {
	return "ServiceInterface"
}

func (p *ServicePlugin) Init(g *Generator) {
	p.Generator = g
}

func (p *ServicePlugin) GeneratorImports(file *FileDescriptor) {
}

func (p.ServicePlugin) Generate(file *FileDescriptor) {
	for _, svc := range file.Service {
		name := CamelCase(*svc.Name)
		p.P("type ", name, " interface {")
		p.In()
		for _, m := range svc.Method {
			method := CamelCase(*m.Name)
			iType := p.ObjectNamed(*m.InputType)
			oType := p.ObjectNamed(*m.OutputType)
			p.P(method, "(in *", p.TypeName(iType), ", out *", p.TypeName(oType), ") error")
		}
		p.Out()
		p.P("}")
	}
}

func init() {
	RegisterPlugin(new(ServicePlugin));
}
