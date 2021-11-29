package db

type Keselamatan struct {
	ID                           int    `json:"id" form:"id"`
	Anti_Lock                    string `json:"anti_lock" form:"anti_lock"`
	Sensor_Parkir                string `json:"sensor_parkir" form:"sensor_parkir"`
	Rearview_Camera              string `json:"rearview_camera" form:"rearview_camera"`
	ElectronicBrake_Distribution string `json:"electronicbrake_distribution" form:"electronicbrake_distribution"`
}
