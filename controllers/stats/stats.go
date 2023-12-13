package metrics

import (
    "os"
    "log"
    "fmt"
    "net/http"

	gin "github.com/gin-gonic/gin"
    dogstatsd "github.com/DataDog/datadog-go/v5/statsd"
)


type Stat struct {
	Metric string   `json:"metric" binding:"required"`
	Value  string   `json:"value binding:"required"`
	Tags []string `json:"tags,omitempty"`
	Type   string `json:"type" binding:"required" default:"count" enums:"count,incr,decr,gauge,set,dist,hist"`
}

var statsd;


func init() {
    STATSD_HOST := os.Getenv("STATSD_HOST");
    STATSD_PORT:= os.Getenv("STATSD_PORT");
    var err error;

    statsd, err = dogstatsd.New(fmt.Sprintf("%s:%s", STATSD_HOST, STATSD_PORT));

    if (err != nil) {
        log.Println("error creating dogstatsd client");
        statsd = nil;
    }
    log.Println("dogstatsd client successfully created");
}


// Submit stat metrics godoc
// @Summary Send stat metrics to the cluster's statsd server 
// @Param metric body string true "metric name"
// @Param value body string true "metric value (float64 if gauge | dist | hist, int64 if count, string otherwise)
// @Param tags body []string false "tags to associate (format: <key>:<val>)" collectionFormat(multi)
// @Param metric_type body string true "metric type" Enums(count, incr, decr, gauge, set, dist, hist) default(count)
// @Schemes
// @Description
// @Tags admin methods
// @Accept json
// @Produce json
// @Success 200
// @Router /stats [post]
// @Security Bearer
func PushStat(c *gin.Context) {
    var stat Stat;
    err := c.ShouldBindJSON(&stat);

    if (err != nil) {
        c.JSON(http.StatusBadRequest, gin.H{ "error" : err.Error});
    }

    // arrange tags
    base_tags := []string{"env:prod", "service:api-gateway"};
    tags := append(base_tags, stat.Tags);

    // cast metric value as needed
    var val_i64 int64;
    var val_f64 float64;
    
    if (stat.Type == "count") {
        val_i64, err = stat.Value.(int64)
    } else if (stat.Type == "gauge" || stat.Type == "dist" || stat.Type == "hist") {
        val_f64, err = stat.Value.(float64)
    }
    
    if (err != nil) {
        c.JSON(http.StatusBadRequest, gin.H{ "error" : err.Error});
    }

    // push metric 
    switch stat.Type {
    case "count":
        statsd.Count(stat.Metric, val_i64, tags, 1);
    case "incr":
        statsd.Incr(stat.Metric, tags, 1);
    case "decr":
        statsd.Decr(stat.Metric, tags, 1);
    case "gauge":
        statsd.Gauge(stat.Metric, val_f64, tags, 1);
    case "set":
        statsd.Set(stat.Metric, stat.Value, tags, 1);
    case "dist":
        statsd.Distribution(stat.Metric, val_f64, tags, 1);
    case "hist":
        statsd.Histogram(stat.Metric, val_f64, tags, 1);
    }

    c.JSON(http.StatusOK, gin.H{"message": "stat pushed"})
}
