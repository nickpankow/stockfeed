package stockfeed

import (
    // "net/url"
    // "net/http"
    // "io/ioutil"
    // "bytes"
    // "encoding/json"
)

type YQL struct {
  Url, Env, Fmt, ApiKey string
}

// map(string)interface{}
func (y *YQL) Query(q string) (string, error){



    return "", nil
}

// Helper function to create queries
func BuildQuery(fields []string, tables []string, where []string) (string) {
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
