package materials

import (
	"fmt"
	"publish/internal/publicate"
)

type Announcement struct {
	Material
	Profile
	Content
}

type Content struct {
	Price       int
	Description string
}

func CreateAnnouncement(name, description, header, Description string, price, ID int, a map[int]publicate.Material) error {
	if err := validateAnnouncement(name, header, ID, price, a); err != nil {
		return err
	}

	announcement := Announcement{
		Profile: Profile{
			Name:        name,
			Description: description,
		},
		Material: Material{
			Header: header,
			Status: "draft",
		},
		Content: Content{
			Price:       price,
			Description: Description,
		},
	}
	announcement.Add("Announcemend created successfully!")
	a[ID] = &announcement

	return nil
}

func (a *Announcement) Edit(name, header, description string, price int) error {
	var message string
	if err := a.validate(header, name, description, price); err != nil {
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

	if description != "" {
		a.Content.Description = description
		message = fmt.Sprint(message, "description ")
	}

	if price != -1 {
		a.Price = price
	}

	message = fmt.Sprint(message, "Edited successfully!")
	a.Add(message)

	return nil
}

func (a *Announcement) Publicate() error {
	if err := a.validatePublicate(); err != nil {
		return err
	}

	a.Status = "published"
	return nil
}
