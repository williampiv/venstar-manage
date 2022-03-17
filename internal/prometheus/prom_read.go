package internalprometheus

// ReadPromTempData reads a temperature value from Prometheus and returns it (currently as a string)
func ReadPromTempData(prometheusIP string, location string) (string, error) {
	// via example: http://${PROM-IP}:9090/api/v1/query?query=temperature_total&location=${LOC}
	// Perhaps take in multiple args like **kwargs after promIP and then &arg=var&arg=var etc etc
	return "", nil
}
