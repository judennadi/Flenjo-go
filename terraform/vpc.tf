# create vpc
resource "aws_vpc" "dev-vpc" {
  cidr_block           = "10.0.0.0/16"
  instance_tenancy     = "default"
  enable_dns_hostnames = true

  tags = {
    "Name" = "dev-vpc"
  }
}

# get an array of availability zones
data "aws_availability_zones" "available-zones" {
}

# create public subnet
resource "aws_subnet" "dev-public-client-subnet" {
  vpc_id                  = aws_vpc.dev-vpc.id
  cidr_block              = "10.0.1.0/24"
  availability_zone       = "us-east-1a"
  map_public_ip_on_launch = true

  tags = {
    "Name" = "dev-public-client-subnet"
  }
}

# create private server subnet
resource "aws_subnet" "dev-private-server-subnet" {
  vpc_id                  = aws_vpc.dev-vpc.id
  cidr_block              = "10.0.2.0/24"
  availability_zone       = "us-east-1a"
  map_public_ip_on_launch = false

  tags = {
    "Name" = "dev-private-server-subnet"
  }
}

# create private data subnet
resource "aws_subnet" "dev-private-data-subnet" {
  vpc_id                  = aws_vpc.dev-vpc.id
  cidr_block              = "10.0.3.0/24"
  availability_zone       = "us-east-1a"
  map_public_ip_on_launch = false

  tags = {
    "Name" = "dev-private-data-subnet"
  }
}

# create private data subnet AZ2
resource "aws_subnet" "dev-private-data-subnet-AZ2" {
  vpc_id                  = aws_vpc.dev-vpc.id
  cidr_block              = "10.0.4.0/24"
  availability_zone       = "us-east-1b"
  map_public_ip_on_launch = false

  tags = {
    "Name" = "dev-private-data-subnet"
  }
}



output "client-public-dns" {
  value = aws_instance.client_ec2.public_dns
}
output "client-public-ip" {
  value = aws_instance.client_ec2.public_ip
}
output "server-private-ip" {
  value = aws_instance.server_ec2.private_ip
}
output "postgres-private-dns" {
  value = aws_db_instance.database_instance.endpoint
}
output "redis-private-dns" {
  value = aws_elasticache_cluster.cache_instance.cache_nodes[0].address
}


