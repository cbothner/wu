/*
* forecast10.go
*
* This file is part of wu.  It contains functions related to
* the -forecast10 switch (10-day forecast).
*
* Written and maintained by Stephen Ramsay <sramsay.unl@gmail.com>
* and Anthony Starks.
*
* Last Modified: Wed Dec 18 16:11:30 CST 2013
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

// printForecast prints the forecast for a given station to standard out
// The dat structure on which it depends is in forecast.go.
func PrintForecast10(obs *Conditions, stationId string) {
  t := obs.Forecast.Txt_forecast
  fmt.Printf("Forecast for %s\n", stationId)
  fmt.Printf("Issued at %s\n", t.Date)
  for _, f := range t.Forecastday {
    fmt.Printf("%s: %s\n", f.Title, f.Fcttext)
  }
}
