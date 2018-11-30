resource "dfd_dfd" "webapp" {
  name = "Webapp"
}

resource "dfd_trust_boundary" "aws_cluster" {
  name   = "Amazon AWS West"
  dfd_id = "${dfd_dfd.webapp.id}"
}

resource "dfd_process" "webserver" {
  name              = "Web Server"
  dfd_id            = "${dfd_dfd.webapp.id}"
  trust_boundary_id = "${dfd_trust_boundary.aws_cluster.id}"
}

resource "dfd_data_store" "logs" {
  name              = "Logs"
  dfd_id            = "${dfd_dfd.webapp.id}"
  trust_boundary_id = "${dfd_trust_boundary.aws_cluster.id}"
}

resource "dfd_flow" "weblogs" {
  name    = "TCP"
  src_id  = "${dfd_process.webserver.id}"
  dest_id = "${dfd_data_store.logs.id}"
}

resource "dfd_trust_boundary" "browser" {
  name   = "External Network"
  dfd_id = "${dfd_dfd.webapp.id}"
}

resource "dfd_process" "client" {
  name              = "Client8"
  dfd_id            = "${dfd_dfd.webapp.id}"
  trust_boundary_id = "${dfd_trust_boundary.browser.id}"
}

resource "dfd_flow" "client-to-web" {
  name    = "HTTPS Request2"
  src_id  = "${dfd_process.client.id}"
  dest_id = "${dfd_process.webserver.id}"
}

resource "dfd_flow" "web-to-client" {
  name    = "HTTPS Response1"
  src_id  = "${dfd_process.webserver.id}"
  dest_id = "${dfd_process.client.id}"
}

resource "dfd_external_service" "google" {
  name   = "Google Analytics"
  dfd_id = "${dfd_dfd.webapp.id}"
}

resource "dfd_flow" "analytics" {
  name    = "HTTPS"
  src_id  = "${dfd_process.client.id}"
  dest_id = "${dfd_external_service.google.id}"
}
