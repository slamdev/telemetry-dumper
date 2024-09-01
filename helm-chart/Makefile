render:
	helm template --namespace=monitoring --debug=false telemetry-dumper .

publish:
ifeq (, $(shell which helm))
	$(error helm binary is not found in path; install it from https://helm.sh/docs/intro/install/)
endif
	helm package --dependency-update --app-version=$${APP_VERSION:-develop} --version=$${CHART_VERSION:-0.0.0-chart-develop} .
	helm push "$$(ls *.tgz)" oci://ghcr.io/slamdev
