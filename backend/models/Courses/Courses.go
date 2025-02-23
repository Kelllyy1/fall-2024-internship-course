/* In this file, I design the Courses schema, which I will use in the future
to create default functions that I will need to interact with each course I
wish to create and store in the database.
*/

/* I want to store the:

Name
Timeline
Deadline
Skills I hope to gain
Area of my career it will impact - which interest am I trying to further

*/

package Courses

import (
	"time"
)

type CourseSchema struct {
	Name         *string    `json:"Name"`
	Timeline     *string    `json:"Timeline"`
	Deadline     *time.Time `json:"Deadline"`
	Skills       *string    `json:"Skills"`
	AreaImpacted *string    `json:"AreaImpacted"`
}

// // Example request:
// {
// 	"Name":"New Course",
// 	"Timeline":"01/01/25",
// 	"Deadline":"01/01/25",
// 	"Skills":"GoLang",
// 	"AreaImpacted":"Backend Development"
// }
