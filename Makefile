



build:
	GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build backend/lambda/main.go;

deploy: build
	zip main.zip main; rm main; \
	mv main.zip builds; \
	cd terraform && terraform apply -auto-approve; \
	cd ../builds; \
	rm main.zip

plan: build
	zip stocksandbonds.zip main; rm main; \
	mv stocksandbonds.zip terraform; \
	cd terraform && terraform init && terraform plan; \
	rm stocksandbonds.zip

clean:
	cd builds; \
	rm main.zip
