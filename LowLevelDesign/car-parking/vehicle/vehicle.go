package vehicle

type VehicleType int

const (
	Small VehicleType = iota
	Medium
	Large
)

type Vehicle struct {
	LicencePlate string
	VehicleType  VehicleType
}
