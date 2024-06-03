package main

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strings"
)
const (
	grafanaUrl = "https://pantheon.grafana.net/api"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	if err := run(log); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}

func run(log *slog.Logger) error {
	// ------------------------------------------------------------------------
	// ENV Variables
	// ------------------------------------------------------------------------
	requiredEnvs := []string{"GRAFANA_TOKEN"}
	errs := ""
	for _, env := range requiredEnvs {
		if os.Getenv(env) == "" {
			errs += fmt.Sprintf("%s, ", env)
		}
	}

	if errs != "" {
		return errors.New(fmt.Sprintf("missing env variables: %s", strings.Trim(errs, " ,")))
	}

	// ------------------------------------------------------------------------
	// CLI Args
	// TODO: Currently required - make this optional
	// ------------------------------------------------------------------------
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < 1 {
		return errors.New("missing args, please specify an api endpoint")
	}
	// ------------------------------------------------------------------------
	// API Magics
	// TODO:
	// 	1. get all dashboards
	// 	2. filter out any dev/scratch/whatever (make a list)
	// 	3. loop through and pull out needed pieces (title, description, tags?)
	// 	4. return all
	// ------------------------------------------------------------------------
  var body string
  var err error
  if argsWithoutProg[0] == "search" {

    queryParam := "%"
    // Default to shwoing ALL dashboards, otherwise do a fuzzy search
    if len(argsWithoutProg) > 1 {
      queryParam = strings.TrimSpace(argsWithoutProg[1])
    }
    body, err = getAllDashboards(queryParam)
    if err != nil {
      return fmt.Errorf("getAllDashboards: %v", err)
    }

  } else {
    body, err = getDashboard(log, argsWithoutProg[0])
    if err != nil {
      return fmt.Errorf("getDashboards [%s]: %v", argsWithoutProg, err)
    }
  }
  fmt.Println(body)

	return nil
}

func getDashboard(log *slog.Logger, dashboardUID string) (string,error) {
	// TODO: DON'T Fix - no one is going to use this in prod ... right?  RIGHT!?
	// CLI Injection RISK!  YEA!!
	url := fmt.Sprintf("%s/dashboards/uid/%s", grafanaUrl, dashboardUID)
	// Inline debugging?  Damn skippy!
	//fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("get request: %v", err)
	}
	bearer := "Bearer " + os.Getenv("GRAFANA_TOKEN")
	req.Header.Add("Authorization", bearer)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("do request: %v", err)
	}
	defer res.Body.Close()
	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		return "", fmt.Errorf("read body: %v", err)
	}

	return string(body), nil
}

func getAllDashboards(queryParam string) (string, error) {

	url := fmt.Sprintf("%s/search?query=%s", grafanaUrl, queryParam)
	// Inline debugging?  Damn skippy!
	//fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("get request: %v", err)
	}
	bearer := "Bearer " + os.Getenv("GRAFANA_TOKEN")
	req.Header.Add("Authorization", bearer)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("do request: %v", err)
	}
	defer res.Body.Close()
	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		return "", fmt.Errorf("read body: %v", err)
	}


  return string(body), nil
}


