package main

type Adminer interface {
	//insert functions admins should be able to perform here
}

type Admin struct {
	//insert data admin should store
}

//implement functions listed in Adminer

type Memberer interface {
	//insert functions regular club members should be able to perform
}

type Member struct {
	//insert data members should be able to store
}

//implement functions listed by Memberer

type Execer interface {
	//insert functions club execs would need
}

type Exec struct {
	//insert data exec users would need to store
}

//implement Execer functions here

type Corporater interface {
	//insert functions company recruiters/speakers would need
}

type Corporate struct {
	//insert data corporate users would need to store
}

//implement Corporater functions here
