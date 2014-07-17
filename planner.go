/*
* history.go
*
* This file is part of wu.  It contains functions related to
* the --planner switch (travel planner based on historical data).
*
* Written and maintained by Stephen Ramsay <sramsay.unl@gmail.com>
* and Anthony Starks.
*
* Last Modified: Wed Dec 18 16:11:55 CST 2013
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
)

type Trip struct {
  Title        string
  Airport_code string
  Error        string
  Chance_of    Chance_of
}

type Chance_of struct {
  Tempoversixty           Tempoversixty
  Chanceofwindyday        Chanceofwindyday
  Chanceofsunnycloudyday  Chanceofsunnycloudyday
  Chanceofprecip          Chanceofprecip
  Chanceofrainday         Chanceofrainday
  Chanceofpartlycloudyday Chanceofpartlycloudyday
  Chanceofthunderday      Chanceofthunderday
  Chanceofhumidday        Chanceofhumidday
  Chanceofcloudyday       Chanceofcloudyday
  Tempoverfreezing        Tempoverfreezing
  Tempoverninety          Tempoverninety
  Chanceoffogday          Chanceoffogday
  Chanceofsnowonground    Chanceofsnowonground
  Chanceoftornadoday      Chanceoftornadoday
  Chanceofsultryday       Chanceofsultryday
  Tempbelowfreezing       Tempbelowfreezing
  Chanceofhailday         Chanceofhailday
  Chanceofsnowday         Chanceofsnowday
}

type Tempoversixty struct {
  Name        string
  Description string
  Percentage  string
}

type Chanceofwindyday struct {
  Name        string
  Description string
  Percentage  string
}

type Chanceofsunnycloudyday struct {
  Name        string
  Description string
  Percentage  string
}

type Chanceofprecip struct {
  Name        string
  Description string
  Percentage  string
}

type Chanceofrainday struct {
  Name        string
  Description string
  Percentage  string
}

type Chanceofpartlycloudyday struct {
  Name        string
  Description string
  Percentage  string
}

type Chanceofthunderday struct {
  Name        string
  Description string
  Percentage  string
}

type Chanceofhumidday struct {
  Name        string
  Description string
  Percentage  string
}

type Chanceofcloudyday struct {
  Name        string
  Description string
  Percentage  string
}

type Tempoverfreezing struct {
  Name        string
  Description string
  Percentage  string
}

type Tempoverninety struct {
  Name        string
  Description string
  Percentage  string
}

type Chanceoffogday struct {
  Name        string
  Description string
  Percentage  string
}

type Chanceofsnowonground struct {
  Name        string
  Description string
  Percentage  string
}

type Chanceoftornadoday struct {
  Name        string
  Description string
  Percentage  string
}

type Chanceofsultryday struct {
  Name        string
  Description string
  Percentage  string
}

type Tempbelowfreezing struct {
  Name        string
  Description string
  Percentage  string
}

type Chanceofhailday struct {
  Name        string
  Description string
  Percentage  string
}

type Chanceofsnowday struct {
  Name        string
  Description string
  Percentage  string
}

func PrintPlanner(obs *Conditions, stationId string) {

  if obs.Trip.Error != "" {
    fmt.Println(obs.Trip.Error)
    os.Exit(0)
  }

  planner := obs.Trip.Chance_of
  fmt.Println(obs.Trip.Title)
  fmt.Println("Station: " + obs.Trip.Airport_code)
  fmt.Println("Chance of: ")
  fmt.Println("   Temps:")
  fmt.Printf("      Over 90 F (32 C): %s%%\n", planner.Tempoverninety.Percentage)
  fmt.Printf("      Between 60 F (15 C) and 90 F (32 C): %s%%\n", planner.Tempoversixty.Percentage)
  fmt.Printf("      Between 32 F (0 C) and 60 (16 C): %s%%\n", planner.Tempoversixty.Percentage)
  fmt.Printf("      Below 32 F (0 C): %s%%\n", planner.Tempbelowfreezing.Percentage)
  fmt.Printf("   Dewpoint above 70 F (21 C): %s%%\n", planner.Chanceofsultryday.Percentage)
  fmt.Printf("   Dewpoint above 60 F (15 C): %s%%\n", planner.Chanceofhumidday.Percentage)
  fmt.Printf("   Winds over 10 mph (15 km/h): %s%%\n", planner.Chanceofwindyday.Percentage)
  fmt.Printf("   %s day: %s%%\n", planner.Chanceofsunnycloudyday.Name, planner.Chanceofsunnycloudyday.Percentage)
  fmt.Printf("   %s day: %s%%\n", planner.Chanceofcloudyday.Name, planner.Chanceofcloudyday.Percentage)
  fmt.Printf("   %s day: %s%%\n", planner.Chanceofpartlycloudyday.Name, planner.Chanceofpartlycloudyday.Percentage)
  fmt.Printf("   %s: %s%%\n", planner.Chanceofprecip.Name, planner.Chanceofprecip.Percentage)
  fmt.Printf("   %s: %s%%\n", planner.Chanceoffogday.Name, planner.Chanceoffogday.Percentage)
  fmt.Printf("   %s: %s%%\n", planner.Chanceofrainday.Name, planner.Chanceofrainday.Percentage)
  fmt.Printf("   %s: %s%%\n", planner.Chanceofthunderday.Name, planner.Chanceofthunderday.Percentage)
  fmt.Printf("   %s: %s%%\n", planner.Chanceoftornadoday.Name, planner.Chanceoftornadoday.Percentage)
  fmt.Printf("   %s: %s%%\n", planner.Chanceofhailday.Name, planner.Chanceofhailday.Percentage)
  fmt.Printf("   %s: %s%%\n", planner.Chanceofsnowday.Name, planner.Chanceofsnowday.Percentage)
  fmt.Printf("   %s: %s%%\n", planner.Chanceofsnowonground.Name, planner.Chanceofsnowonground.Percentage)
}
