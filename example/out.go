package models




type TicketType int


const (
  TicketModerationRequestPermissions TicketType = iota + 1
  TicketOfferModeration
)



type Ticket struct {
	tableName struct{} `sql:"tickets,alias:t" pg:",discard_unknown_columns"`


  Type     TicketType `xorm:"type" sql:"type" json:"type"`
  Title     string `xorm:"title" sql:"title" json:"title"`
  Text     string `xorm:"text" sql:"text" json:"text"`
  ObjectID     int `xorm:"objectId" sql:"objectId" json:"objectId"`
  ID     int `json:"id" xorm:"pk autoincr 'ticketId'" sql:"ticketId"`
  CreatedAt     time.Time `xorm:"createdAt" sql:"createdAt" json:"createdAt"`
  UpdatedAt     time.Time `xorm:"updatedAt" sql:"updatedAt" json:"updatedAt"`
  Params     struct{} `sql:"params" json:"params" xorm:"jsonb notnull default '{}'"`
}

func (t *Ticket) TableName() string {
  return "tickets"
}





