provider "aws" {
  region = "us-east-1"
}

resource "aws_instance" "app_server" {
  ami           = "ami-0182f373e66f89c85" # Example Amazon Linux AMI
  instance_type = "t2.micro"

  tags = {
    Name = "AppServer"
  }
}
