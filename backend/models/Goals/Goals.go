/* In this file, I design the Goals schema, which I will use in the future
to create default functions that I will need to interact with each course I
wish to create and store in the database.
*/

/* I want to store the

Nickname
Timeline
Deadline
Description
Resource that can help me achieve this goal

*/

package Goals

import (
	"time"
)

type Goals struct {
	Nickname     string    `json:"Name"`
	Timeline     string    `json:"Timeline"`
	Deadline     time.Time `json:"Deadline"`
	Skills       string    `json:"Skills"`
	AreaImpacted string    `json:"AreaImpacted"`
}
