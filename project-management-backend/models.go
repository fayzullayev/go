package main

type Task struct {
	id    uint32
	title string
}

type Project struct {
	id          uint32
	title       string
	description string
	dueDate     string
}

func getAllProjects() ([]Project, error) {

	query := `SELECT * FROM projects;`

	var projects []Project = []Project{}

	rows, err := DB.Query(query)

	if err != nil {
		return nil, err
	}
	rows.Close()

	for rows.Next() {
		var project Project

		err = rows.Scan(&project.id, &project.title, &project.description, &project.dueDate)
		if err != nil {
			return nil, err
		}

		projects = append(projects, project)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return projects, nil

}
