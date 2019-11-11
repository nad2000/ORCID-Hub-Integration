locals {
    UPSTREAM_URL = "${aws_api_gateway_deployment.ORCIDHUB_INTEGRATION_API_deployment.invoke_url}${aws_api_gateway_resource.ORCIDHUB_INTEGRATION_API_Resource_Call.path}"
}

resource "null_resource" "orcidhub_webhook" {
  depends_on = [aws_api_gateway_deployment.ORCIDHUB_INTEGRATION_API_deployment]
  provisioner "local-exec" {
    command = "./create.sh"
    environment = {
			CLIENT_ID = "${local.CLIENT_ID}"
			CLIENT_SECRET = "${local.CLIENT_SECRET}"
			ENV = "${local.ENV}"
			APIKEY = "${local.APIKEY}"
			UPSTREAM_URL = "${local.UPSTREAM_URL}"
    }
	}
  provisioner "local-exec" {
		when    = "destroy"
    command = "./destroy.sh"
    environment = {
			CLIENT_ID = "${local.CLIENT_ID}"
			CLIENT_SECRET = "${local.CLIENT_SECRET}"
			ENV = "${local.ENV}"
    }
	}
}

output "UPSTREAM_URL" {
	value = "${local.UPSTREAM_URL}"
}
