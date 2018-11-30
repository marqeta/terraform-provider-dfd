[![Build Status](https://cloud.drone.io/api/badges/marqeta/terraform-provider-dfd/status.svg)](https://cloud.drone.io/marqeta/terraform-provider-dfd)

# terraform-provider-dfd

A Terraform provider for generating Data Flow Diagrams in DOT (Graphviz) format.

## Requirements

* Terraform ~> 0.11

## Installing the Provider

Visit the [Releases Page](https://github.com/marqeta/terraform-provider-dfd/releases)
and download the appropriate archive for your system.

```sh
# Extract the provider
$> cd /path/to/workdir
$> echo "provider \"dfd\" {}" > main.tf
$> tar -xzvf terraform-provider-dfd_v0.0.1_darwin-amd64.tar.gz
$> terraform init

Initializing provider plugins...

Terraform has been successfully initialized!

You may now begin working with Terraform. Try running "terraform plan" to see
any changes that are required for your infrastructure. All Terraform commands
should now work.

If you ever set or change modules or backend configuration for Terraform,
rerun this command to reinitialize your working directory. If you forget, other
commands will detect it and remind you to do so if necessary.
```

## Configuring the Provider

The provider can be configured to point to a specific `.dot` file. If you do
not specify this value, one will be chosen for you by default and stored in
your current working directory.

```hcl
provider "dfd" {
  dot_path = "/path/to/my/dfd/dfd.dot"
}
```

## Using the Provider

This provider has the following resources:

* dfd_dfd - The Data Flow Diagram you wish to create
* dfd_data_store - An element of the DFD, such as a database or file store
* dfd_flow - A transmission of data from one entity to another, such as an HTTPS connection
* dfd_external_service - A third party service in use by the application you are modeling
* dfd_process - A managed service or actor within an application, such as a user or a web server
* dfd_trust_boundary - A collection of elements within a Trust Boundary

Please see the `examples/` directory for usage information.

## Generating Graphs

The current representation of your Data Flow Diagram (based on the outcome of
the last `terraform apply` action) is stored in the `.dot` file you've
configured. The following command assumes you've installed Graphviz to your
system and that the `dot` executable is available in your `$PATH`

```sh
$> dot -Tpng dfd.dot -o dfd.png # generate a PNG file
$> dot -Tpdf dfd.dot -o dfd.png # generate a PDF file
```

You can also copy the contents of your `.dot` file and paste it into a tool
like (GraphvizOnline)[https://dreampuf.github.io/GraphvizOnline].

## Troubleshooting

This provider is lightweight. When you believe there is something wrong, the
easiest thing to do is to remove both the Terraform state files and the `.dot`
file generated via a command like `rm terraform.tfstate* dfd.dot`. You can then
re-run `terraform apply` to regenerate everything.
