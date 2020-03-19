package helpers

const (
	PLAYER=1
	GROUNDOWNER=2
	TRAINER=3
)

var roles_array [3]int;

func Roles_checker(role int) bool{
	roles_array[0] = 1;
	roles_array[1] = 2;
	roles_array[2] = 3;
	for _, value := range roles_array{
		if value == role{
			return true
		}
	}
	return false
}