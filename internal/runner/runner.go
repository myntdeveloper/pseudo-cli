package runner

import (
	"fmt"

	"github.com/myntdeveloper/pseudo-cli/internal/models"
	"github.com/myntdeveloper/pseudo-cli/internal/store"
)

type Runner struct {
	Store *store.Store
}

func (r *Runner) Save(pseudonym models.Pseudonym) error {
	return r.Store.Save(pseudonym.Name, pseudonym.Command, pseudonym.Description, pseudonym.Tag)
}

func (r *Runner) List(tag string) error {
	items, err := r.Store.List(tag)
	if err != nil {
		return err
	}
	for _, item := range items {
		fmt.Printf("%s: %s", item.Name, item.Command)
		if item.Description != "" {
			fmt.Printf(" — %s", item.Description)
		}
		if item.Tag != "" {
			fmt.Printf(" [%s]", item.Tag)
		}
		fmt.Println()
	}
	return nil
}

func (r *Runner) Show(name string) error {
	p, err := r.Store.GetByName(name)
	if err != nil {
		return err
	}
	fmt.Printf("Name: %s\nCommand: %s\n", p.Name, p.Command)
	if p.Description != "" {
		fmt.Printf("Description: %s\n", p.Description)
	}
	if p.Tag != "" {
		fmt.Printf("Tag: %s\n", p.Tag)
	}
	return nil
}

func (r *Runner) Remove(name string) error {
	err := r.Store.Delete(name)
	return err
}
