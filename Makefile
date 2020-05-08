



build:
	go build backend/lambda/main.go;

deploy: build
	zip stocksandbonds.zip main; rm main; \
	mv stocksandbonds.zip terraform; \
	cd terraform && terraform apply -auto-approve; \
	rm stocksandbonds.zip

plan: build
	zip stocksandbonds.zip main; rm main; \
	mv stocksandbonds.zip terraform; \
	cd terraform && terraform init && terraform plan; \
	rm stocksandbonds.zip


