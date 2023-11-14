package bilf

type Settings struct {
	Database string
}
type Tag struct {
	Id   int
	Name string
}
type Daily struct {
	Tag []int
}
