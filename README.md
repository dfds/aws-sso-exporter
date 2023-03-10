# aws-sso-exporter

A Prometheus exporter that gathers data via AWS CloudTrail(Authenticate and Federate events), calculates metrics based on that data, and exports it in a format supported by Prometheus.

## Installation

Container images based on the master branch is available on [Docker Hub](https://hub.docker.com/r/dfdsdk/aws-sso-exporter) via the *latest* tag.

No pre-compiled binaries are provided. A simple 'go build cmds/server.go' should do the trick.

Examples of Kubernetes manifests can be found in the *k8s/* directory.

## Configuration

### Environment variables

- **ASE_AWS_REGION** : "eu-west-1" // Set the AWS region where *aws-sso-exporter* should look for CloudTrail events
- **ASE_WORKERINTERVAL** : "600" // Set the sleep interval between scrape jobs

### AWS

IAM permissions needed:

- cloudtrail:LookupEvents

## Todo

- Make *startTime* in aws/cloudtrail.go configurable
- Create Helm chart
