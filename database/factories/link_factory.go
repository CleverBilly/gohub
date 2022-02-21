package factories

import (
	"github.com/bxcodec/faker/v3"
	"gohub/app/models/link"
)

func MakeLinks(times int) []link.Link {

	var objs []link.Link

	for i := 0; i < times; i++ {
		model := link.Link{
			Name: faker.Username(),
			URL:  faker.URL(),
		}
		objs = append(objs, model)
	}

	return objs
}
