package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
	"time"
)

type Todo struct {
	Text        string    `json:"todo"`
	Done        bool      `json:"completed"`
	CreatedAt   time.Time `json:"createdAt"`
	CompletedAt time.Time `json:"completedAt"`
}

type Todos []Todo

func (t *Todos) Add(todo string) {

	task := Todo{
		Text:        todo,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, task)
}

func (t *Todos) Complete(index int) error {
	ls := *t

	if index <= 0 || index > len(ls) {
		return errors.New("Invalid index")
	}

	ls[index-1].CompletedAt = time.Now()
	ls[index-1].Done = true

	return nil
}

func (t *Todos) Delete(index int) error {
	ls := *t

	if index <= 0 || index > len(ls) {
		return errors.New("Invalid index")
	}

	*t = append(ls[:index-1], ls[index:]...)

	return nil
}

func (t *Todos) Load(f string) error {

	file, err := os.ReadFile(f)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return errors.New("File not found")
	}

	if len(file) == 0 {
		return err
	}

	err = json.Unmarshal(file, t)
	if err != nil {
		return errors.New("Filed to marshal data")
	}

	return nil
}

func (t *Todos) Store(f string) error {

	data, err := json.Marshal(t)
	if err != nil {
		return errors.New("Faild to marshal data")
	}
	return os.WriteFile(f, data, 0644)
}

func (t *Todos) PrintTodos() {

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.AlignRight)

	// Print the header
	fmt.Fprintln(w, "#\tTask\tDone\tCreated\tCompleted\t")

	for index, todo := range *t {
		index++
		fmt.Fprintf(w, "%v\t%v\t%v\t%v\t %v\n", index, todo.Text, todo.Done, todo.CreatedAt.Format(time.RFC822), todo.CompletedAt.Format(time.RFC822))
	}

	err := w.Flush()
	if err != nil {
		log.Fatal(err.Error())
	}
}
