/*
* wu - a small, fast command-line application for retrieving weather
* data from Weather Underground
*
* Main and associated functions.
*
* Written and maintained by Stephen Ramsay <sramsay.unl@gmail.com>
* and Anthony Starks.
*
* Last Modified: Sun Sep 01 16:38:59 CDT 2013
*
* Copyright © 2010-2013 by Stephen Ramsay and Anthony Starks.
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
  "encoding/json"
  "flag"
  "fmt"
  "io/ioutil"
  "net/http"
  "os"
  "regexp"
  "strings"
)

type Config struct {
  Key     string
  Station string
}

var (
  help         bool
  version      bool
  doall        bool
  doalmanac    bool
  doalerts     bool
  doconditions bool
  dolookup     bool
  doforecast   bool
  doforecast10 bool
  doastro      bool
  doyesterday  bool
  dotides      bool
  dohistory    string
  doplanner    string
  date         string
  conf         Config
)

// Struct common to several data streams
type Date struct {
  Pretty string
  Hour   string
  Min    string
  Mon    string
  Mday   string
  Year   string
}

const defaultStation = "KLNK"

// GetVersion returns the version of the package
func GetVersion() string {
  return "3.9.7"
}

// GetConf returns the API key and weather station from
// the configuration file at $HOME/.condrc
func ReadConf() {

  if b, err := ioutil.ReadFile(os.Getenv("HOME") + "/.condrc"); err == nil {
    jsonErr := json.Unmarshal(b, &conf)
    CheckError(jsonErr)
  } else {
    fmt.Println("You must create a .condrc file in $HOME.")
    os.Exit(0)
  }
}

// Options handles commandline options and returns a 
// possibly updated weather station string
func Options() string {

  var station, sconf string

  if conf.Station == "" {
    sconf = defaultStation
  } else {
    sconf = conf.Station
  }

  flag.BoolVar(&doconditions, "conditions", false, "Reports the current weather conditions")
  flag.BoolVar(&doalerts, "alerts", false, "Reports any active weather alerts")
  flag.BoolVar(&dolookup, "lookup", false, "Lookup the codes for the weather stations in a particular area")
  flag.BoolVar(&doastro, "astro", false, "Reports sunrise, sunset, and lunar phase")
  flag.BoolVar(&doforecast, "forecast", false, "Reports the current (3-day) forecast")
  flag.BoolVar(&doforecast10, "forecast10", false, "Reports the current (7-day) forecast")
  flag.BoolVar(&doalmanac, "almanac", false, "Reports average high, low and record temperatures")
  flag.BoolVar(&doyesterday, "yesterday", false, "Reports yesterday's weather data")
  flag.StringVar(&dohistory, "history", "", "Reports historical data for a particular day --history=\"YYYYMMDD\"")
  flag.StringVar(&doplanner, "planner", "", "Reports historical data for a particular date range (30-day max) --planner=\"MMDDMMDD\"")
  flag.BoolVar(&dotides, "tides", false, "Reports tidal data (if available")
  flag.BoolVar(&help, "help", false, "Print this message")
  flag.BoolVar(&version, "version", false, "Print the version number")
  flag.BoolVar(&doall, "all", false, "Show all weather data")
  flag.StringVar(&station, "s", sconf,
    "Weather station: \"city, state-abbreviation\", (US or Canadian) zipcode, 3- or 4-letter airport code, or LAT,LONG")
  flag.Parse()

  // Check for correct usage of wu -lookup
  if dolookup {
    if len(os.Args) == 3 {
      station = os.Args[len(os.Args)-1]
    } else {
      fmt.Println("Usage: wu -lookup [station] where station is a \"city, state-abbreviation\", (US or Canadian) zipcode, 3- or 4-letter airport code, or LAT,LONG")
      os.Exit(0)
    }
  }

  if help {
    flag.PrintDefaults()
    os.Exit(0)
  }

  if version {
    fmt.Println("Wu " + GetVersion())
    fmt.Println("Copyright 2010-2014 by Stephen Ramsay and")
    fmt.Println("Anthony Starks. Data courtesy of Weather")
    fmt.Println("Underground, Inc. is subject to Weather")
    fmt.Println("Underground Data Feed Terms of Service.")
    fmt.Println("The program itself is free software, and")
    fmt.Println("you are welcome to redistribute it under")
    fmt.Println("certain conditions.  See LICENSE for details.")
    os.Exit(0)
  }

  // Trap for city-state combinations (e.g. "San Francisco, CA") and
  // make them URL-friendly (e.g. "CA/SanFranciso")
  cityStatePattern := regexp.MustCompile("([A-Za-z ]+), ([A-Za-z ]+)")

  if cityState := cityStatePattern.FindStringSubmatch(station); cityState != nil {
    station = cityState[2] + "/" + cityState[1]
    station = strings.Replace(station, " ", "_", -1)
  }
  return station
}

// BuildURL returns the URL required by the Weather Underground API
// from the query type, station id, and API key
func BuildURL(infoTypes []string, stationId string) string {

  const URLstem = "http://api.wunderground.com/api/"
  const query = "/q/"
  const format = ".json"

  var URL string

  for i, value := range infoTypes {
    if value == "history" {
      infoTypes[i] += "_" + dohistory
    } else if value == "planner" {
      infoTypes[i] += "_" + doplanner
    }
  }
  URL = URLstem + conf.Key + "/" + strings.Join(infoTypes, "/") + query + stationId + format

   //fmt.Println(URL) //DEBUG

  return URL
}

// Fetch does URL processing
func Fetch(url string) ([]byte, error) {
//fmt.Println("Calling API") //DEBUG

  res, err := http.Get(url)
  CheckError(err)
  if res.StatusCode != 200 {
    fmt.Fprintf(os.Stderr, "Bad HTTP Status: %d\n", res.StatusCode)
    return nil, err
  }
  b, err := ioutil.ReadAll(res.Body)
  res.Body.Close()
  return b, err
}

// CheckError exits on error with a message
func CheckError(err error) {
  if err != nil {
    fmt.Fprintf(os.Stderr, "Fatal error\n%v\n", err)
    os.Exit(1)
  }
}

func init() {
  ReadConf()
}

type Conditions struct {
  Alerts []Alerts
  Almanac Almanac
  Current_observation Current
  Forecast Forecast
  History History
  Location SLocation
  Moon_phase Moon_phase
  Sunrise Sunrise
  Sunset Sunset
  Tide Tide
  Trip Trip
}

// weather prints various weather information for a specified station
func weather(operations []string, station string) {
  url := BuildURL(operations, station)
  b, err := Fetch(url)
  CheckError(err)

  var obs Conditions
  jsonErr := json.Unmarshal(b, &obs)
  CheckError(jsonErr)
  for _, operation := range operations {
    operation = strings.Split(operation, "_")[0]
    switch operation {
    case "almanac":
      PrintAlmanac(&obs, station)
    case "astronomy":
      PrintAstro(&obs, station)
    case "alerts":
      PrintAlerts(&obs, station)
    case "conditions":
      PrintConditions(&obs)
    case "forecast":
      PrintForecast(&obs, station)
    case "forecast10day":
      PrintForecast10(&obs, station)
    case "yesterday":
      PrintHistory(&obs, station)
    case "history":
      PrintHistory(&obs, station)
    case "planner":
      PrintPlanner(&obs, station)
    case "tide":
      PrintTides(&obs, station)
    case "geolookup":
      PrintLookup(&obs)
    }
  }
}

func main() {
  stationId := Options()
  operations := make([]string, 0)
  if dohistory != "" && doplanner != "" {
    fmt.Println(
      "Weather Underground does not support making a history\n" +
      "request and a planner request at the same time.")
    os.Exit(1)
  }
  if doall {
    operations = append(operations,"conditions")
    operations = append(operations,"forecast")
    operations = append(operations,"forecast10day")
    operations = append(operations,"alerts")
    operations = append(operations,"almanac")
    operations = append(operations,"history")
    operations = append(operations,"planner")
    operations = append(operations,"yesterday")
    operations = append(operations,"astronomy")
    operations = append(operations,"tide")
    operations = append(operations,"geolookup")
  }
  if doalerts {
    operations = append(operations,"alerts")
  }
  if doalmanac {
    operations = append(operations,"almanac")
  }
  if doastro {
    operations = append(operations,"astronomy")
  }
  if doconditions {
    operations = append(operations,"conditions")
  }
  if doforecast {
    operations = append(operations,"forecast")
  }
  if doforecast10 {
    operations = append(operations,"forecast10day")
  }
  if dohistory != "" {
    operations = append(operations,"history")
  }
  if doyesterday {
    operations = append(operations,"yesterday")
  }
  if doplanner != "" {
    operations = append(operations,"planner")
  }
  if dotides {
    operations = append(operations,"tide")
  }
  if dolookup {
    operations = append(operations,"geolookup")
  }
  if flag.NFlag() == 0 {
    operations = append(operations,"conditions")
  }
  weather(operations, stationId)
}
