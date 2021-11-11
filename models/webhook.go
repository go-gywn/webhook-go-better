package models

type Webhook struct {
	LabelMapper      map[string]string `yaml:"labelMapper" default:"{}"`
	AnnotationMapper map[string]string `yaml:"annotationMapper" default:"{}"`
	Targets          map[string]Target `yaml:"targets" default:"{}"`
}

type Target struct {
	Template string `yaml:"template" default:"template.tpl"`
	API      string
	Params   string
	Method   string
	Type     string
}

func (t *Target) AddTarget() error {
	return nil
}

func (t *Target) DelTarget() error {
	return nil
}
