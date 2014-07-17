/*
* tide.go
*
* This file is part of wu.  It contains functions related to
* the -tides switch (high and low tide data).
*
* Written and maintained by Stephen Ramsay <sramsay.unl@gmail.com>
* and Anthony Starks.
*
* Last Modified: Wed Dec 18 16:12:03 CST 2013
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
  "os"
  "strconv"
  "time"
)

type Tide struct {
  Tideinfo    []Tideinfo
  Tidesummary []Tidesummary
}

type Tideinfo struct {
  Tidesite string
}

type Tidesummary struct {
  Date Date // Defined in wu.go
  Data Data
}

type Data struct {
  Height string
  Type   string
}

// printTides prints the tidal data for given station to standard out
func PrintTides(obs *Conditions, stationID string) {
  tide := obs.Tide
  info := tide.Tideinfo
  summary := tide.Tidesummary

  if len(summary) == 0 {
    fmt.Println("No tidal data available.")
    os.Exit(0)
  }

  fmt.Printf("Tidal data for %s\n", info[0].Tidesite)

  var date_string string
  var prev_date string

  for _, s := range summary {
    month, _ := strconv.Atoi(s.Date.Mon)
    hour, _ := strconv.Atoi(s.Date.Hour)
    prev_date = date_string
    date_string = time.Month(month).String() + " " + s.Date.Mday + ", " + s.Date.Year + ":"
    if date_string != prev_date {
      fmt.Println(date_string)
    }
    if hour < 13 {
      fmt.Printf("     %s at %d:%s AM\n", s.Data.Type, hour, s.Date.Min)
    } else {
      hour = hour - 12
      fmt.Printf("     %s at %d:%s PM\n", s.Data.Type, hour, s.Date.Min)
    }
  }
}
