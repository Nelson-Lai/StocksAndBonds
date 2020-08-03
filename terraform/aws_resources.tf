provider "aws" {
  region     = "us-east-2"
}

resource "aws_lambda_function" "stocksandbonds" {
  filename      = "../builds/main.zip"
  function_name = "stocksandbonds"
  role          = "arn:aws:iam::120364169916:role/service-role/lambda_basic_execution"
  handler       = "main"
  runtime = "go1.x"
  timeout = 30
  source_code_hash = filebase64sha256("../builds/main.zip")
}



