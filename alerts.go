/*
* alerts.go
*
* This file is part of wu.  It contains functions related to
* the -alerts switch (active weather alerts).
*
* Written and maintained by Stephen Ramsay <sramsay.unl@gmail.com>
* and Anthony Starks.
*
* Last Modified: Wed Dec 18 16:10:55 CST 2013
*
* Copyright © 2010-2014 by Stephen Ramsay and Anthony Starks.
*
* wu is free software; you can redistribute it and/or modify
* it under the terms of the GNU General Public License as published by
* the Free Software Foundation; either version 3, or (at your option)
* any later version.
*
* wu is distributed in the hope that it will be useful, but WITHOUT
* ANY WARRANTY; without even the implied warranty of MERCHANTABILITY
* or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public
* License for more details.
*
* You should have received a copy of the GNU General Public License
* along with wu; see the file COPYING.  If not see
* <http://www.gnu.org/licenses/>.
 */

package main

import (
  "fmt"
)

type Alerts struct {
  Date        string
  Expires     string
  Description string
  Message     string
}

// printAlerts prints the alerts for a given station to standard out
func PrintAlerts(obs *Conditions, stationId string) {
  if len(obs.Alerts) == 0 {
    fmt.Println("No active alerts")
  } else {
    fmt.Printf("Station: %s\n", stationId)
    for _, a := range obs.Alerts {
      fmt.Printf("### %s ###\n\nIssued at %s\nExpires at %s\n%s\n",
        a.Description, a.Date, a.Expires, a.Message)
    }
  }
}
