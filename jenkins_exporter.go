package main
import (
	"fmt"
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)
var (
	target = target.rstrip("/")
	user = user
	password = password
	insecure = insecure
)
type jenkinsCollector struct {
	entries *prometheus.Desc
}
func (*jenkinsCollector) collect () {
	
}
func (*jenkinsCollector) request_data () {
	
}
func (func request_data) parsejobs (myurl) {
	
}
func (*jenkinsCollector) setup_empty_prometheus_metrics () {
	
}


func init() {
	registerCollector("jenkins", defaultEnabled, NewJenkinsCollector)
}

// NewJenkinCollector returns a new Collector exposing Jenkins stats.
func NewJenkinsControler()(Collector, error) {
	return &jenkinsCollector{
		entries: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "jenkins", "entries"),
			"Jenkins entries created",
			[]string{"device"}, nil,
		),
	}, nil
}

func getJenkinsEntries() (map[string]uint32, error){
	...
	return entries, nil
}

//需要更改细节
func parse_args() {
	parser = argparse.ArgumentParser(
		description='jenkins exporter args jenkins address and port'
	)
	parser.add_argument(
		'-j', '--jenkins',
		metavar='jenkins',
		required=False,
		help='server url from the jenkins api',
		default=os.environ.get('JENKINS_SERVER', 'http://jenkins:8080')
	)
	parser.add_argument(
		'--user',
		metavar='user',
		required=False,
		help='jenkins api user',
		default=os.environ.get('JENKINS_USER')
	)
	parser.add_argument(
		'--password',
		metavar='password',
		required=False,
		help='jenkins api password',
		default=os.environ.get('JENKINS_PASSWORD')
	)
	parser.add_argument(
		'-p', '--port',
		metavar='port',
		required=False,
		type=int,
		help='Listen to this port',
		default=int(os.environ.get('VIRTUAL_PORT', '9118'))
	)
	parser.add_argument(
		'-k', '--insecure',
		dest='insecure',
		required=False,
		action='store_true',
		help='Allow connection to insecure Jenkins API',
		default=False
	)
	return parser.parse_args()
}


func Update() {
	entries, err := getJenkinsEntries()
	if err != nil {
		return fmt.Errorf("could not get Jenkins entries: %s", err)
	}

	for device, entryCount := range entries {
		ch <- prometheus.MustNewConstMetric(
			c.entries, prometheus.GaugeValue, float64(entryCount), device)
	}

	return nil
}
func main() {
	return
}