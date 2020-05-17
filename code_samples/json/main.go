// Binary prom2csvdate generates CSV values returned by a prometheus query with date range.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/prometheus/common/model"
)

// Results holds nested JSON response from prometheus.
type Results struct {
	Status string `json:"status"`
	Data   QueryResult
}

// QueryResult holds results of a prometheus query.
type QueryResult struct {
	ResultType model.ValueType `json:"resultType"`
	Value      []model.Sample  `json:"result"`
}

var (
	//query        = flag.String("query", "", "PromQL query")
	//server       = flag.String("server", "http://localhost:9090/", "Prometheus server address")
	header = flag.Bool("header", false, "Whether to print CSV header, default: false")
	//startTime    = flag.String("startTime", "", "Input Start time in this format 2020-05-10T20:11:00.781Z")
	//endTime      = flag.String("endTime", "", "Input End time in this format 2020-05-10T20:11:00.781Z")
	//timeInterval = flag.String("timeInterval", "", "Input time interval or step size in this format 1h")
)

func main() {
	js := `{"status":"success","data":{"resultType":"matrix","result":[{"metric":{"__name__":"ap_client_count","ap_mac":"0062ec4a5720","ap_name":"-INTEST-SVL-BLR-01-1562I-002","entity":"lab","hotspot":"UNKNOWN","job":"cisco_wlc_snmp","sys_mac":"0062ec066740","vendor":"cisco"},"values":[[1589141430.781,"1"],[1589145030.781,"1"],[1589148630.781,"1"]]},{"metric":{"__name__":"ap_client_count","ap_mac":"502fa81258f8","ap_name":"-INTEST-SVL-BLR-01-1562I-001","entity":"lab","hotspot":"UNKNOWN","job":"cisco_wlc_snmp","sys_mac":"bc26c7536820","vendor":"cisco"},"values":[[1589141430.781,"0"],[1589145030.781,"0"],[1589148630.781,"0"]]},{"metric":{"__name__":"ap_client_count","ap_mac":"706bb9705420","ap_name":"-STAGING-SVL-SWF-01-1832-001","entity":"lab","hotspot":"SWF","job":"cisco_wlc_snmp","sys_mac":"706bb971b980","vendor":"cisco"},"values":[[1589141430.781,"1"],[1589145030.781,"1"],[1589148630.781,"1"]]},{"metric":{"__name__":"ap_client_count","ap_name":"-BRTEST-GRU-LNK-AP-02-R500-001","entity":"lab","hotspot":"LNK","instance":"hele.default.svc:9990","job":"hele","vendor":"ruckus"},"values":[[1589141430.781,"0"],[1589145030.781,"0"],[1589148630.781,"0"]]},{"metric":{"__name__":"ap_client_count","ap_name":"-STAGING-SVL-SIT-AP-02-R320-001","entity":"lab","hotspot":"SIT","instance":"hele.default.svc:9990","job":"hele","vendor":"ruckus"},"values":[[1589141430.781,"0"],[1589145030.781,"0"],[1589148630.781,"0"]]},{"metric":{"__name__":"ap_client_count","ap_name":"-ZATEST-SVL-THINK-AP-02-R310-001","entity":"lab","hotspot":"THINK","instance":"hele.default.svc:9990","job":"hele","vendor":"ruckus"},"values":[[1589141430.781,"1"],[1589145030.781,"1"],[1589148630.781,"1"]]}]}}`
	flag.Parse()
	/*
		if *query == "" {
			fmt.Println("Usage:")
			flag.PrintDefaults()
			os.Exit(1)
		}
		if *startTime == "" || *endTime == "" || *timeInterval == "" {
			fmt.Println("Usage:")
			flag.PrintDefaults()
			os.Exit(1)
		}

			response, err := http.Get(*server + "/api/v1/query?query=" + url.QueryEscape(*query) + "&" + *startTime + "&" + *endTime + "&" + timeInterval)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			defer response.Body.Close()

			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
	*/
	res := &Results{}
	if err := json.Unmarshal([]byte(js), res); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	// Collect unique labelnames from result
	var uniqueLabels = make(map[string]bool)

	for _, sample := range res.Data.Value {
		for key := range sample.Metric {
			k := string(key)
			// skip the metric name itself
			if strings.Contains(k, model.MetricNameLabel) {
				continue
			}
			uniqueLabels[k] = true
		}
	}

	// Sort labelnames
	var sortedLabels []string
	for key := range uniqueLabels {
		sortedLabels = append(sortedLabels, key)
	}
	sort.Strings(sortedLabels)

	// Print CSV header ?
	if *header {
		fmt.Println(strings.Join(sortedLabels, ",") + ",value")
	}

	// Print labelvalues and metric value for each sorted label
	for _, sample := range res.Data.Value {
		var row []string
		for _, l := range sortedLabels {
			row = append(row, fmt.Sprintf("%s", sample.Metric[model.LabelName(l)]))
		}
		// Print the labels and metric value
		fmt.Println(strings.Join(row, ",") + fmt.Sprintf(",%v", sample.Value))
	}

}
