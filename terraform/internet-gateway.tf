# create internet gateway
resource "aws_internet_gateway" "dev-igw" {
  vpc_id = aws_vpc.dev-vpc.id

  tags = {
    Name = "dev-igw"
  }
}

# create public route table
resource "aws_route_table" "dev-public-rt" {
  vpc_id = aws_vpc.dev-vpc.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.dev-igw.id
  }

  tags = {
    Name = "dev-public-rt"
  }
}

# associate public subnets to public route table
resource "aws_route_table_association" "public-subnet-association" {
  route_table_id = aws_route_table.dev-public-rt.id
  subnet_id      = aws_subnet.dev-public-client-subnet.id
}
