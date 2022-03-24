package controllers

import (
	"log"
	"net/http"
	"todo_app/app/models"
)

func top(w http.ResponseWriter, r *http.Request) {
	// t, err := template.ParseFiles("app/views/templates/top.html")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// t.Execute(w, "Hello")

	_, err := session(w, r)
	if err != nil {
		generateHTML(w, "Hello", "layout", "public_navbar", "top")
	} else {
		http.Redirect(w, r, "/todos", 302)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	session, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", 302)
	} else {
		user, err := session.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		todos, _ := user.GetTodosByUser()
		user.Todos = todos
		generateHTML(w, user, "layout", "private_navbar", "index")
	}
}

func todoNew(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "todo_new")
	}
}

func todoSave(w http.ResponseWriter, r *http.Request) {
	session, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		content := r.PostFormValue("content")
		user, err := session.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		if err := user.CreateTodo(content); err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/todos", 302)
	}
}

func todoEdit(w http.ResponseWriter, r *http.Request, id int) {
	session, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err := session.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		todo, err := models.GetTodo(id)
		if err != nil {
			log.Println(err)
		}
		if r.Method == "GET" {
			generateHTML(w, todo, "layout", "private_navbar", "todo_edit")
		} else if r.Method == "POST" {
			todoUpdate(w, r, todo)
		}
	}
}

func todoUpdate(w http.ResponseWriter, r *http.Request, todo models.Todo) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	todo.Content = r.PostFormValue("content")
	if err := todo.UpdateTodo(); err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/todos", 302)
}

func todoDelete(w http.ResponseWriter, r *http.Request, id int) {
	session, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err := session.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		todo, err := models.GetTodo(id)
		if err != nil {
			log.Println(err)
		}
		if err := todo.DeleteTodo(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/todos", 302)
	}
}
