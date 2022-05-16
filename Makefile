version = 99.0.0
provider_path = registry.terraform.io/xaque208/tunnelbroker/$(version)/linux_amd64/

install_linux:
	go build -o terraform-provider-tunnelbroker_$(version)

	mkdir -p ~/.terraform.d/plugins//$(provider_path)
	mv terraform-provider-tunnelbroker_$(version)  ~/.terraform.d/plugins/$(provider_path)
