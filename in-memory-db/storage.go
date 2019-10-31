package inMemoryDb

var database map[int]Processed

func store()  {
	cl:=ClientPro{
		ClientID: "125 - fseopfnskldmvdpi",
		ConnectedAt: "2018-10-25",
		Keepalive: 4,
		Node: "123@fgfg",
	}
	var cla []ClientPro
	cla=append(cla, cl)
	se:=SessionsPro{
		CreatedAt: "2018-10-25",
		ClientID: "125 - fseopfnskldmvdpi",
		Subscriptions: 2,
	}
	var sea []SessionsPro
	sea= append(sea, se)
	su:=SubscriptionsPro{
		ClientID: "125 - fseopfnskldmvdpi",
		Qos: 5,
		Topic: "testing",
	}
	var sua []SubscriptionsPro
	sua= append(sua, su)
	value:=Processed{
		cla,
		sea,
		sua,
	}
	database[20]=value
	database[25]=value
	database[30]=value
}