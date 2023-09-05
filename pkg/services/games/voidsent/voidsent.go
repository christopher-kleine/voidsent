package voidsent

type Voidsent struct {
	sessions []string
	roles    byte
	password string
	ownerID  string
}

func New(password string, ownerID string, roles byte) *Voidsent {
	return &Voidsent{
		sessions: []string{
			ownerID,
		},
		roles:    roles,
		password: password,
		ownerID:  ownerID,
	}
}

func (v *Voidsent) Start() {
	//go func() {
	//	for msg := range v.in {
	//		fmt.Printf("Log: %+v\n", msg)
	//	}
	//}()
}
