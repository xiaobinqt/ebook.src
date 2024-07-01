package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://hq1.appsflyer.com/api/raw-data/export/app/club.nota/installs_report/v5?from=2023-07-17&to=2023-07-17"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "text/csv")
	req.Header.Add("authorization", "Bearer eyJhbGciOiJBMjU2S1ciLCJjdHkiOiJKV1QiLCJlbmMiOiJBMjU2R0NNIiwidHlwIjoiSldUIiwiemlwIjoiREVGIn0.YTS6XuTkjulvDUjI_Dlj0DWr9Qd5B9FxThT8Jumo8L-wZEPKm99JqQ.5NMsoE0Soa010Kw7.OiNH7J62YSOE6Nxe3d9N7QyAzxHypsT6OVvLD8BzgP6abFcke9KVSXxlAeftr095hsJ6SBVkENZg5j6tQpj3nNwinMTA0M9fLD_hc2m2iIECmbOLF1eLlTxPjLT3_d5ORHhOXnZBySJzs-8AwP9NpxVDT4ByYuQPyJGRDfNo-xiHRJd7vslRSQ-z1A2gFNlaN_peC0c_GjZNitF48JLTqc_sOwrpVEOjrYf1ikEkNN1oXezu7vyQwG8XEPfMyXXGACIucqsH3Q-SNJFp04VpwaXsH1ofGTPa-nyT40F-WW2TOaxqTlteLjPHrgbAjXPQybjVkkXRSjlxVENS8HzIrG3l2qLm4gP2YTL4r4HXb-2yAEMJZEqHfFGHEyemdRD8-QpC7L-SZ4lT9Utvp2xk40BWi6LvsZDfAcoQj39bIJbuYtOvwGAfoi19A27Mqdgg2lctgZVcFsnkraXlvTCyyAfZ2Az7Z29M66ipdEBraFaFjUb3YWO76wZFJVcxOwKnacf-BtT2Z31IZ3ZaPSOjh2k.3ya1Tqp6eBAX9Gsdxp5PmQ")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("do err", err.Error())
		return
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(body))
}
