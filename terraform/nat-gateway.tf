# allocate elastic ip that will be used for the nat-gateway
resource "aws_eip" "eip1" {
  vpc = true

  tags = {
    Name = "dev-eip1"
  }
}

# create nat gateway
resource "aws_nat_gateway" "dev-ngw" {
  allocation_id = aws_eip.eip1.id
  subnet_id     = aws_subnet.dev-public-client-subnet.id

  tags = {
    "Name" = "dev-ngw"
  }

  depends_on = [aws_internet_gateway.dev-igw]
}

# create private route table
resource "aws_route_table" "dev-private-rt" {
  vpc_id = aws_vpc.dev-vpc.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_nat_gateway.dev-ngw.id
  }

  tags = {
    Name = "dev-private-rt"
  }
}


# associate private server subnets to private route table
resource "aws_route_table_association" "private-server-subnet-association" {
  route_table_id = aws_route_table.dev-private-rt.id
  subnet_id      = aws_subnet.dev-private-server-subnet.id
}
# associate private data subnets to private route table
resource "aws_route_table_association" "private-data-subnet-association" {
  route_table_id = aws_route_table.dev-private-rt.id
  subnet_id      = aws_subnet.dev-private-data-subnet.id
}
