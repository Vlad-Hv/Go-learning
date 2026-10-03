package materials

import (
	"fmt"
	"publish/internal/publicate"
)

type Article struct {
	Profile
	Material
	Context string
}

func CreateArticle(name, description, header, context string, ID int, articles map[int]publicate.Material) error {
	err := validateArticles(ID, name, header, articles)
	if err != nil {
		return err
	}
	article := Article{
		Profile: Profile{
			Name:        name,
			Description: description,
		},
		Material: Material{
			Header: header,
			Status: "draft",
		},
		Context: context,
	}

	article.Add("Article created succesfully")
	articles[ID] = &article

	return nil
}

func (a *Article) Edit(name, header, context string, price int) error {
	var message string
	if err := a.validate(name, header, context); err != nil {
		return err
	}

	if name != "" {
		a.Name = name
		message = fmt.Sprint(message, "name ")
	}

	if header != "" {
		a.Header = header
		message = fmt.Sprint(message, "header ")
	}

	if context != "" {
		a.Context = context
		message = fmt.Sprint(message, "context ")
	}

	message = fmt.Sprint(message, "Edited succesfully")
	a.Add(message)
	return nil
}

func (a *Article) Publicate() error {
	if err := a.validatePublicate(); err != nil {
		return err
	}

	a.Status = "published"
	return nil
}
