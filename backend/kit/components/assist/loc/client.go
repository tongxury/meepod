package loc

type Client struct {
	locs []*Loc
}

func NewLocClient() *Client {
	c := &Client{
		locs: get(),
	}

	return c
}

func (c *Client) Get() []*Loc {
	return c.locs
}

func get() []*Loc {

	var rsp []*Loc

	for _, province := range provinces {

		yp := &Loc{
			Name: province.Name,
			Id:   province.Id,
		}

		for _, city := range cities[province.Id] {

			yc := &Loc{
				Name: city.Name,
				Id:   city.Id,
			}

			for _, area := range areas[city.Id] {
				yc.Children = append(yc.Children, &Loc{
					Name: area.Name,
					Id:   area.Id,
				})
			}

			yp.Children = append(yp.Children, yc)

		}

		rsp = append(rsp, yp)
	}

	return rsp
}
