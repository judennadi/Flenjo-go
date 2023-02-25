# create security group for client
resource "aws_security_group" "client-security-group" {
  name   = "dev-client-sg"
  vpc_id = aws_vpc.dev-vpc.id

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = -1
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    "Name" = "dev-client-sg"
  }
}

# create security group for server
resource "aws_security_group" "server-security-group" {
  name   = "dev-server-sg"
  vpc_id = aws_vpc.dev-vpc.id

  ingress {
    from_port       = 5000
    to_port         = 5000
    protocol        = "tcp"
    security_groups = [aws_security_group.client-security-group.id]
  }
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = -1
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    "Name" = "dev-server-sg"
  }
}

# create security group for database
resource "aws_security_group" "database-security-group" {
  name   = "dev-database-sg"
  vpc_id = aws_vpc.dev-vpc.id

  ingress {
    from_port       = 5432
    to_port         = 5432
    protocol        = "tcp"
    security_groups = [aws_security_group.server-security-group.id]
  }
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = -1
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    "Name" = "dev-database-sg"
  }
}

# create security group for cache
resource "aws_security_group" "cache-security-group" {
  name   = "dev-cache-sg"
  vpc_id = aws_vpc.dev-vpc.id

  ingress {
    from_port       = 6379
    to_port         = 6379
    protocol        = "tcp"
    security_groups = [aws_security_group.client-security-group.id]
  }
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = -1
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    "Name" = "dev-cache-sg"
  }
}

# create security group for ssh
resource "aws_security_group" "ssh-security-group" {
  name   = "dev-ssh-sg"
  vpc_id = aws_vpc.dev-vpc.id

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = -1
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    "Name" = "dev-ssh-sg"
  }
}
