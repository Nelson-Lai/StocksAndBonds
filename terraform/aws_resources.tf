provider "aws" {
  region     = "us-east-2"
}

resource "aws_lambda_function" "stocksandbonds" {
  filename      = "stocksandbonds.zip"
  function_name = "stocksandbonds"
  role          = "arn:aws:iam::120364169916:role/service-role/lambda_basic_execution"
  handler       = "main"
  runtime = "go1.x"
  timeout = 30
}



