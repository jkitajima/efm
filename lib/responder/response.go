package responder

type MetaField struct {
	Meta MetaObject `json:"meta"`
}

type MetaObject struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func NewMetaField(status int, msg string) *MetaField {
	return &MetaField{
		Meta: MetaObject{
			Status:  status,
			Message: msg,
		},
	}
}

type DataField struct {
	Data any `json:"data"`
}
