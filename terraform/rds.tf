# create database subnet group
resource "aws_db_subnet_group" "database-subnet-group" {
  name       = "dev-database-subnets"
  subnet_ids = [aws_subnet.dev-private-data-subnet.id, aws_subnet.dev-private-data-subnet-AZ2.id]

  tags = {
    Name = "dev-database-subnets"
  }
}

# launch an rds instance from a database snapshot
resource "aws_db_instance" "database_instance" {
  allocated_storage      = 10
  storage_type           = "gp2"
  db_name                = var.db_name
  engine                 = "postgres"
  identifier             = "dev-db"
  instance_class         = "db.t3.micro"
  username               = var.db_username
  password               = var.db_password
  skip_final_snapshot    = true
  availability_zone      = "us-east-1a"
  multi_az               = false
  db_subnet_group_name   = aws_db_subnet_group.database-subnet-group.name
  vpc_security_group_ids = [aws_security_group.database-security-group.id]
}

