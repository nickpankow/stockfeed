package stockfeed

import (
    "net/url"
    "net/http"
    "io/ioutil"
    "bytes"
    "encoding/json"
)

const YQL_URL = "https://query.yahooapis.com/v1/public/yql?"
const YQL_ENV = "env=http://datatables.org/alltables.env"
const YQL_FMT = "format=json"

type QueryResponse struct {
    created, lang string
    results map[string]interface{}

}

func buildQuery(fields []string, tables []string, where []string) (string) {
    // Validate
    if len(fields) == 0 || len(tables) == 0 || len(where) == 0{
        return ""
    }
    
    // Setup Buffer
    var query_buffer bytes.Buffer
    
    // Select
    query_buffer.WriteString("select ")
    for key, value := range fields{
        if key > 0 {
            query_buffer.WriteString(",")
        }
        query_buffer.WriteString(value) 
    }
    
    // From
    query_buffer.WriteString(" from ")
    for key, value := range tables{
        if key > 0 {
            query_buffer.WriteString(",")
        }
        query_buffer.WriteString(value) 
    }
    
    // Where
    query_buffer.WriteString(" where ")
    for key, value := range where{
        if key > 0 {
            query_buffer.WriteString(" AND ")
        }
        query_buffer.WriteString(value) 
    }
    
    return query_buffer.String()
}

func buildQueryURL(query string) (string){
    return YQL_URL + "q=" + url.QueryEscape(query) + "&" + YQL_FMT + "&" + YQL_ENV
}

func query(yql string) (map[string]interface{}, error){
    query_url := buildQueryURL(yql)
    // fmt.Print(query_url,"\n")
    resp, err := http.Get(query_url)
    
    // HTTP error
    if err != nil {
        return nil, err
    }
    
    resp_str, err := ioutil.ReadAll(resp.Body)
    // Read error
    if err != nil {
        return nil, err
    }
    // fmt.Printf("Resp: %s", resp_str)
    
    defer resp.Body.Close()
    
    // Parse JSON response
    // var json_array map[string]string
    var json_array map[string]interface{}
    // json_string = json.NewDecoder(strings.newReader(string(resp_str)))
    // json_string := string(resp_str)
    err = json.Unmarshal(resp_str, &json_array)
    
    return json_array, nil
    // return string(resp_str), nil
}